package user

import (
	"otter/api/common"
	"otter/constants/api"
	"otter/constants/userstatus"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/po/rolePo"
	"otter/po/userPo"
	"otter/service/jwt"
	"otter/service/sha3"

	"github.com/valyala/fasthttp"
)

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(signUp SignUpReqVo) error {
	run := func() interface{} {
		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)
		columnValues := map[string]interface{}{
			userPo.Acc:  signUp.Acc,
			userPo.Pwd:  encryptPwd,
			userPo.Name: signUp.Name,
		}

		if _, err := mysql.Insert(userPo.Table, columnValues); err != nil {
			return err
		}

		return nil
	}

	return jobqueue.User.NewSignUpJob(run).(error)
}

// SignIn dao
func (dao *Dao) SignIn(ctx *fasthttp.RequestCtx, signIn SignInReqVo) (*SignInResVo, api.RespStatus, error) {
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
		return nil, mysql.ErrMsgHandler(err), err
	}

	// check pwd
	if entity.Pwd != sha3.Encrypt(signIn.Pwd) {
		return nil, api.DataError, nil
	}

	// check account status
	if entity.Status != string(userstatus.Active) {
		return nil, api.AccInactive, nil
	}

	var signInResVo SignInResVo
	token, _ := jwt.Generate(entity.ID, entity.Acc, entity.Name, entity.RoleCode, roleEnt.Name)
	signInResVo = SignInResVo{
		Token: token,
	}
	return &signInResVo, api.Success, nil
}

// Update dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, updateData UpdateReqVo) error {
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
	args = append(args, updateData.ID)

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
		return err
	}

	return nil
}

// List dao
func (dao *Dao) List(ctx *fasthttp.RequestCtx, listReqVo ListReqVo) (common.PageRespVo, error) {
	args := []interface{}{}

	var whereSQL string
	if listReqVo.Active == "true" {
		whereSQL = "WHERE " + userPo.Status + "=?"
		args = append(args, userstatus.Active)
	}

	page := mysql.Page{
		TableName:   userPo.Table,
		ColumnNames: []string{userPo.ID, userPo.Acc, userPo.Name, userPo.RoleCode, userPo.Status},
		UniqueKey:   userPo.PK,
		OrderBy:     userPo.ID,
		Page:        listReqVo.Page,
		Limit:       listReqVo.Limit,
	}

	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
		Total:   0,
	}
	err := mysql.QueryPage(page, whereSQL, args, func(result mysql.Rows, total int) error {
		rows := result.Rows
		for rows.Next() {
			var record ListResVo
			err := rows.Scan(&record.ID, &record.Acc, &record.Name, &record.RoleCode, &record.Status)
			if err != nil {
				return err
			}
			list.Records = append(list.Records, record)
		}

		list.Total = total
		return nil
	})
	if err != nil {
		return list, err
	}

	return list, nil
}
