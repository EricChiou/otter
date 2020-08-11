package user

import (
	"fmt"

	"otter/api/common"
	"otter/api/role"
	"otter/constants/api"
	"otter/constants/userstatus"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/service/apihandler"
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
		defer func() {
			wait <- 1
		}()

		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)

		var entity Entity
		kv := map[string]interface{}{
			entity.Col().Acc:  signUp.Acc,
			entity.Col().Pwd:  encryptPwd,
			entity.Col().Name: signUp.Name,
		}
		_, err := mysql.Insert(entity.Table(), kv)
		if err != nil {
			fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
			return
		}

		fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, nil, nil))
	})
	<-wait
}

// SignIn dao
func (dao *Dao) SignIn(ctx *fasthttp.RequestCtx, signIn SignInReqVo) {
	var entity Entity
	var roleEnt role.Entity
	sql := "SELECT user.#idCol, user.#accCol, user.#pwdCol, user.#nameCol, user.#roleCodeCol, user.#statusCol, role.#roleNameCol " +
		"FROM #userT user " +
		"INNER JOIN #roleT role ON user.#roleCodeCol=role.#codeCol " +
		"WHERE user.#accCol=?"
	var param map[string]string = make(map[string]string)
	param["userT"] = entity.Table()
	param["idCol"] = entity.Col().ID
	param["accCol"] = entity.Col().Acc
	param["pwdCol"] = entity.Col().Pwd
	param["nameCol"] = entity.Col().Name
	param["roleCodeCol"] = entity.Col().RoleCode
	param["statusCol"] = entity.Col().Status
	param["roleT"] = roleEnt.Table()
	param["roleNameCol"] = roleEnt.Col().Name
	param["codeCol"] = roleEnt.Col().Code

	args := []interface{}{signIn.Acc}
	err := mysql.QueryRow(sql, param, args, func(result mysql.RowResult) error {
		row := result.Row
		err := row.Scan(&entity.ID, &entity.Acc, &entity.Pwd, &entity.Name, &entity.RoleCode, &entity.Status, &roleEnt.Name)
		if err != nil {
			return err
		}

		return nil
	})
	// check account existing
	if err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	// check pwd
	if entity.Pwd != sha3.Encrypt(signIn.Pwd) {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.DataError, nil, nil))
		return
	}

	// check account status
	if entity.Status != string(userstatus.Active) {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.AccInactive, nil, nil))
		return
	}

	var response SignInResVo
	token, _ := jwt.Generate(entity.ID, entity.Acc, entity.Name, entity.RoleCode, roleEnt.Name)
	response = SignInResVo{
		Token: token,
	}
	fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, response, nil))
}

// Update dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, payload jwt.Payload, updateData UpdateReqVo) {
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

	_, err := mysql.Update(entity.Table(), set, where)
	if err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, nil, nil))
}

// List dao
func (dao *Dao) List(ctx *fasthttp.RequestCtx, listReqVo ListReqVo) {
	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
		Total:   0,
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
	if listReqVo.Active == "true" {
		where[entity.Col().Status] = userstatus.Active
	}
	orderBy := entity.Col().ID
	total, err := mysql.Page(entity.Table(), entity.PK(), column, where, orderBy, listReqVo.Page, listReqVo.Limit, func(result mysql.RowsResult) error {
		rows := result.Rows
		for rows.Next() {
			var record ListResVo
			err := rows.Scan(&record.ID, &record.Acc, &record.Name, &record.RoleCode, &record.Status)
			if err != nil {
				return err
			}
			list.Records = append(list.Records, record)
		}

		return nil
	})
	if err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), list, err))
		return
	}
	list.Total = total

	fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, list, nil))
}
