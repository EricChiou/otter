package user

import (
	"fmt"

	"otter/api/common"
	cons "otter/constants"
	"otter/db/mysql"
	api "otter/service/apihandler"
	"otter/service/jwt"
	"otter/service/sha3"

	"github.com/valyala/fasthttp"
)

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(ctx *fasthttp.RequestCtx, signUp SignUpReqVo) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, nil, err))
		return
	}

	// encrypt password
	encryptPwd := sha3.Encrypt(signUp.Pwd)

	var entity Entity
	kv := map[string]interface{}{
		entity.Col().Email: signUp.Email,
		entity.Col().Pwd:   encryptPwd,
		entity.Col().Name:  signUp.Name,
	}
	_, err = mysql.Insert(tx, entity.Table(), kv)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, nil, nil))
}

// SignIn dao
func (dao *Dao) SignIn(ctx *fasthttp.RequestCtx, signIn SignInReqVo) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()

	var entity Entity
	var response SignInResVo
	column := []string{
		entity.Col().ID,
		entity.Col().Email,
		entity.Col().Pwd,
		entity.Col().Name,
		entity.Col().Role,
	}
	where := map[string]interface{}{
		entity.Col().Email:  signIn.Email,
		entity.Col().Active: true,
	}
	row := mysql.QueryRow(tx, entity.Table(), column, where)
	err = row.Scan(&entity.ID, &entity.Email, &entity.Pwd, &entity.Name, &entity.Role)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDataError, response, err))
		return
	}

	if entity.Pwd != sha3.Encrypt(signIn.Pwd) {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDataError, response, nil))
		return
	}

	token, _ := jwt.Generate(entity.ID, entity.Email, entity.Name, entity.Role)
	response = SignInResVo{
		Token: token,
	}
	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, response, nil))
}

// Update dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, payload jwt.Payload, updateData UpdateReqVo) (string, interface{}) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.RSDBError, err
	}

	var entity Entity
	set := map[string]interface{}{}
	if len(updateData.Name) != 0 {
		set[entity.Col().Name] = updateData.Name
	}
	if len(updateData.Pwd) != 0 {
		set[entity.Col().Pwd] = sha3.Encrypt(updateData.Pwd)
	}
	var where map[string]interface{} = make(map[string]interface{})
	if updateData.ID != 0 {
		where[entity.Col().ID] = updateData.ID
	} else {
		where[entity.Col().ID] = payload.ID
	}

	_, err = mysql.Update(tx, entity.Table(), set, where)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.RSSuccess, nil
}

// List dao
func (dao *Dao) List(ctx *fasthttp.RequestCtx, page, limit int, active bool) (common.PageRespVo, string, interface{}) {
	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    page,
		Limit:   limit,
	}
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return list, cons.RSDBError, err
	}

	var entity Entity
	column := []string{
		entity.Col().ID,
		entity.Col().Email,
		entity.Col().Name,
		entity.Col().Role,
		entity.Col().Active,
	}
	where := map[string]interface{}{}
	if active {
		where[entity.Col().Active] = true
	}
	orderBy := entity.Col().ID
	rows, err := mysql.Page(tx, entity.Table(), entity.PK(), column, where, orderBy, page, limit)
	defer rows.Close()
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}

	for rows.Next() {
		var data ListDataVo
		err = rows.Scan(&data.ID, &data.Email, &data.Name, &data.Identity, &data.Active)
		if err != nil {
			return list, mysql.ErrMsgHandler(err), err
		}
		list.Records = append(list.Records, data)
	}

	var total int
	err = tx.QueryRow("SELECT COUNT(*) FROM " + entity.Table()).Scan(&total)
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}
	list.Total = total

	return list, cons.RSSuccess, nil
}
