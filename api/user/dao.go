package user

import (
	"fmt"

	"otter/api/common"
	"otter/api/role"
	cons "otter/constants"
	"otter/db/mysql"
	"otter/jobqueue"
	api "otter/service/apihandler"
	"otter/service/jwt"
	"otter/service/sha3"

	"github.com/valyala/fasthttp"
)

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(ctx *fasthttp.RequestCtx, signUp SignUpReqVo) {
	wait := make(chan int)
	jobqueue.User.SignUp.Add(func() {
		tx, err := mysql.DB.Begin()
		defer func() {
			tx.Commit()
			wait <- 1
		}()
		if err != nil {
			fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, nil, err))
			return
		}

		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)

		var entity Entity
		kv := map[string]interface{}{
			entity.Col().Acc:  signUp.Acc,
			entity.Col().Pwd:  encryptPwd,
			entity.Col().Name: signUp.Name,
		}
		_, err = mysql.Insert(tx, entity.Table(), kv)
		if err != nil {
			fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
			return
		}

		fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, nil, nil))
	})
	<-wait
}

// SignIn dao
func (dao *Dao) SignIn(ctx *fasthttp.RequestCtx, signIn SignInReqVo) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()

	var entity Entity
	var roleEnt role.Entity
	var response SignInResVo

	selectStr := "SELECT t1." + entity.Col().ID + ", t1." + entity.Col().Acc + ", t1." + entity.Col().Pwd + ", t1." + entity.Col().Name + ", t1." + entity.Col().RoleCode + ", t2." + roleEnt.Col().Name + " "
	from := "FROM `" + entity.Table() + "` as t1 "
	join := "INNER JOIN `" + roleEnt.Table() + "` as t2 "
	on := "ON t1." + entity.Col().RoleCode + "=t2." + roleEnt.Col().Code + " "
	where := "WHERE t1." + entity.Col().Acc + "=? "
	and := "AND t1." + entity.Col().Status + "=?"
	row := tx.QueryRow(selectStr+from+join+on+where+and,
		signIn.Acc,
		cons.UserStstus.Active,
	)
	err = row.Scan(&entity.ID, &entity.Acc, &entity.Pwd, &entity.Name, &entity.RoleCode, &roleEnt.Name)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDataError, response, err))
		return
	}

	if entity.Pwd != sha3.Encrypt(signIn.Pwd) {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDataError, response, nil))
		return
	}

	token, _ := jwt.Generate(entity.ID, entity.Acc, entity.Name, entity.RoleCode, roleEnt.Name)
	response = SignInResVo{
		Token: token,
	}
	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, response, nil))
}

// Update dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, payload jwt.Payload, updateData UpdateReqVo) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, nil, err))
		return
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
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, nil, nil))
}

// List dao
func (dao *Dao) List(ctx *fasthttp.RequestCtx, page, limit int, active bool) {
	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    page,
		Limit:   limit,
	}
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, list, err))
		return
	}

	var entity Entity
	column := []string{
		entity.Col().ID,
		entity.Col().Acc,
		entity.Col().Name,
		entity.Col().RoleCode,
		entity.Col().Status,
	}
	where := map[string]interface{}{}
	if active {
		where[entity.Col().Status] = cons.UserStstus.Active
	}
	orderBy := entity.Col().ID
	rows, err := mysql.Page(tx, entity.Table(), entity.PK(), column, where, orderBy, page, limit)
	defer rows.Close()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), list, err))
		return
	}

	for rows.Next() {
		var data ListDataVo
		err = rows.Scan(&data.ID, &data.Acc, &data.Name, &data.RoleCode, &data.Status)
		if err != nil {
			fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), list, err))
			return
		}
		list.Records = append(list.Records, data)
	}

	var total int
	err = tx.QueryRow("SELECT COUNT(*) FROM " + entity.Table()).Scan(&total)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), list, err))
		return
	}
	list.Total = total

	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, list, nil))
}
