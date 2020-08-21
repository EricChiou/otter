package user

import (
	"otter/api/common"
	"otter/api/role"
	"otter/constants/api"
	"otter/constants/userstatus"
	"otter/db/mysql"
	"otter/jobqueue/queues"
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
	queues.User.SignUp.Add(func() {
		defer func() {
			wait <- 1
		}()

		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)

		var entity Entity
		kvParams := mysql.GetSQLParamsInstance()
		kvParams.Add(entity.Col().Acc, signUp.Acc)
		kvParams.Add(entity.Col().Pwd, encryptPwd)
		kvParams.Add(entity.Col().Name, signUp.Name)

		_, err := mysql.Insert(entity.Table(), kvParams)
		if err != nil {
			apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
			return
		}

		apihandler.Response(ctx, api.Success, nil, nil)
	})
	<-wait
}

// SignIn dao
func (dao *Dao) SignIn(ctx *fasthttp.RequestCtx, signIn SignInReqVo) {
	var entity Entity
	var roleEnt role.Entity
	sql := ""
	sql += "SELECT user.#idCol, user.#accCol, user.#pwdCol, user.#nameCol, user.#roleCodeCol, user.#statusCol, role.#roleNameCol "
	sql += "FROM #userT user "
	sql += "INNER JOIN #roleT role ON user.#roleCodeCol=role.#codeCol "
	sql += "WHERE user.#accCol=:acc"
	param := mysql.GetSQLParamsInstance()
	param.Add("userT", entity.Table())
	param.Add("idCol", entity.Col().ID)
	param.Add("accCol", entity.Col().Acc)
	param.Add("pwdCol", entity.Col().Pwd)
	param.Add("nameCol", entity.Col().Name)
	param.Add("roleCodeCol", entity.Col().RoleCode)
	param.Add("statusCol", entity.Col().Status)
	param.Add("roleT", roleEnt.Table())
	param.Add("roleNameCol", roleEnt.Col().Name)
	param.Add("codeCol", roleEnt.Col().Code)
	param.Add("acc", signIn.Acc)

	err := mysql.QueryRow(sql, param, func(result mysql.RowResult) error {
		row := result.Row
		err := row.Scan(&entity.ID, &entity.Acc, &entity.Pwd, &entity.Name, &entity.RoleCode, &entity.Status, &roleEnt.Name)
		if err != nil {
			return err
		}

		return nil
	})
	// check account existing
	if err != nil {
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
		return
	}

	// check pwd
	if entity.Pwd != sha3.Encrypt(signIn.Pwd) {
		apihandler.Response(ctx, api.DataError, nil, nil)
		return
	}

	// check account status
	if entity.Status != string(userstatus.Active) {
		apihandler.Response(ctx, api.AccInactive, nil, nil)
		return
	}

	var response SignInResVo
	token, _ := jwt.Generate(entity.ID, entity.Acc, entity.Name, entity.RoleCode, roleEnt.Name)
	response = SignInResVo{
		Token: token,
	}
	apihandler.Response(ctx, api.Success, response, nil)
}

// Update dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, payload jwt.Payload, updateData UpdateReqVo) {
	var entity Entity
	setParams := mysql.GetSQLParamsInstance()
	if len(updateData.Name) != 0 {
		setParams.Add(entity.Col().Name, updateData.Name)
	}
	if len(updateData.Pwd) != 0 {
		setParams.Add(entity.Col().Pwd, sha3.Encrypt(updateData.Pwd))
	}

	whereParams := mysql.GetSQLParamsInstance()
	if updateData.ID != 0 {
		whereParams.Add(entity.Col().ID, updateData.ID)
	} else {
		whereParams.Add(entity.Col().ID, payload.ID)
	}

	_, err := mysql.Update(entity.Table(), setParams, whereParams)
	if err != nil {
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
		return
	}

	apihandler.Response(ctx, api.Success, nil, nil)
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

	sql := ""
	sql += "SELECT "

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
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), list, err)
		return
	}
	list.Total = total

	apihandler.Response(ctx, api.Success, list, nil)
}
