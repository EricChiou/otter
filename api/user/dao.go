package user

import (
	"otter/api/common"
	"otter/constants/api"
	"otter/constants/userstatus"
	"otter/db/mysql"
	"otter/jobqueue/queues"
	"otter/po/rolePo"
	"otter/po/userPo"
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
		columnValues := map[string]interface{}{
			userPo.Acc:  signUp.Acc,
			userPo.Pwd:  encryptPwd,
			userPo.Name: signUp.Name,
		}

		_, err := mysql.Insert(userPo.Table, columnValues)
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
	var entity userPo.Entity
	var roleEnt rolePo.Entity

	args := []interface{}{signIn.Acc}

	param := mysql.SQLParamsInstance()
	param.Add("user", userPo.Table)
	param.Add("id", userPo.ID)
	param.Add("acc", userPo.Acc)
	param.Add("pwd", userPo.Pwd)
	param.Add("name", userPo.Name)
	param.Add("roleCode", userPo.RoleCode)
	param.Add("status", userPo.Status)
	param.Add("role", rolePo.Table)
	param.Add("roleName", rolePo.Name)
	param.Add("code", rolePo.Code)

	sql := "SELECT user.#id, user.#acc, user.#pwd, user.#name, user.#roleCode, user.#status, role.#roleName "
	sql += "FROM #user user "
	sql += "INNER JOIN #role role ON user.#roleCode=role.#code "
	sql += "WHERE user.#acc=?"

	err := mysql.QueryRow(sql, param, args, func(result mysql.Row) error {
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
	args := []interface{}{}
	var setSQL string
	if len(updateData.Name) != 0 {
		setSQL += ", #name=?"
		args = append(args, updateData.Name)
	}
	if len(updateData.Pwd) != 0 {
		setSQL += ", #pwd=?"
		args = append(args, sha3.Encrypt(updateData.Pwd))
	}
	setSQL = setSQL[2:] + " "

	if updateData.ID != 0 {
		args = append(args, updateData.ID)
	} else {
		args = append(args, payload.ID)
	}

	params := mysql.SQLParamsInstance()
	params.Add("user", userPo.Table)
	params.Add("name", userPo.Name)
	params.Add("pwd", userPo.Pwd)
	params.Add("id", userPo.ID)

	sql := "UPDATE #user "
	sql += "SET " + setSQL
	sql += "WHERE #id=?"

	_, err := mysql.Exec(sql, params, args)
	if err != nil {
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
		return
	}

	apihandler.Response(ctx, api.Success, nil, nil)
}

// List dao
func (dao *Dao) List(ctx *fasthttp.RequestCtx, listReqVo ListReqVo) {
	args := []interface{}{(listReqVo.Page - 1) * listReqVo.Limit, listReqVo.Limit}
	whereArgs := []interface{}{}

	var whereSQL string
	if listReqVo.Active == "true" {
		whereSQL = "WHERE #status=?"
		whereArgs = append(whereArgs, userstatus.Active)
	}

	params := mysql.SQLParamsInstance()
	params.Add("user", userPo.Table)
	params.Add("pk", userPo.PK)
	params.Add("id", userPo.ID)
	params.Add("acc", userPo.Acc)
	params.Add("name", userPo.Name)
	params.Add("roleCode", userPo.RoleCode)
	params.Add("status", userPo.Status)

	sql := "SELECT #id, #acc, #name, #roleCode, #status "
	sql += "FROM #user "
	sql += "    JOIN ( "
	sql += "    SELECT #pk FROM #user "
	sql += "    ORDER BY #id "
	sql += "    LIMIT ?, ? "
	sql += ") t "
	sql += "USING ( #pk ) "
	sql += whereSQL

	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
		Total:   0,
	}
	err := mysql.Query(sql, params, append(args, whereArgs...), func(result mysql.Rows) error {
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
		apihandler.ResponsePage(ctx, mysql.ErrMsgHandler(err), list, err)
		return
	}

	sql = "SELECT COUNT(*) FROM #user " + whereSQL

	var total int
	err = mysql.QueryRow(sql, params, whereArgs, func(result mysql.Row) error {
		return result.Row.Scan(&total)
	})
	if err != nil {
		apihandler.ResponsePage(ctx, mysql.ErrMsgHandler(err), list, err)
		return
	}
	list.Total = total

	apihandler.ResponsePage(ctx, api.Success, list, nil)
}
