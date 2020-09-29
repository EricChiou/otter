package user

import (
	"otter/api/common"
	"otter/bo/userbo"
	"otter/constants/userstatus"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/po/rolepo"
	"otter/po/userpo"
	"otter/service/sha3"
)

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(signUp SignUpReqVo) error {
	run := func() interface{} {
		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)
		columnValues := map[string]interface{}{
			userpo.Acc:  signUp.Acc,
			userpo.Pwd:  encryptPwd,
			userpo.Name: signUp.Name,
		}

		if _, err := mysql.Insert(userpo.Table, columnValues); err != nil {
			return err
		}

		return nil
	}

	return jobqueue.User.NewSignUpJob(run).(error)
}

// SignIn dao
func (dao *Dao) SignIn(signInReqVo SignInReqVo) (userbo.SignInBo, error) {
	var signInBo userbo.SignInBo

	args := []interface{}{signInReqVo.Acc}

	param := mysql.SQLParamsInstance()
	param.Add("user", userpo.Table)
	param.Add("id", userpo.ID)
	param.Add("acc", userpo.Acc)
	param.Add("pwd", userpo.Pwd)
	param.Add("name", userpo.Name)
	param.Add("roleCode", userpo.RoleCode)
	param.Add("status", userpo.Status)
	param.Add("role", rolepo.Table)
	param.Add("roleName", rolepo.Name)
	param.Add("code", rolepo.Code)

	sql := "SELECT user.#id, user.#acc, user.#pwd, user.#name, user.#roleCode, user.#status, role.#roleName "
	sql += "FROM #user user "
	sql += "INNER JOIN #role role ON user.#roleCode=role.#code "
	sql += "WHERE user.#acc=?"

	err := mysql.QueryRow(sql, param, args, func(result mysql.Row) error {
		row := result.Row
		err := row.Scan(&signInBo.ID, &signInBo.Acc, &signInBo.Pwd, &signInBo.Name, &signInBo.RoleCode, &signInBo.Status, &signInBo.RoleName)
		if err != nil {
			return err
		}

		return nil
	})
	// check account existing
	if err != nil {
		return signInBo, err
	}

	return signInBo, nil
}

// Update dao
func (dao *Dao) Update(updateData UpdateReqVo) error {
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
	params.Add("user", userpo.Table)
	params.Add("name", userpo.Name)
	params.Add("pwd", userpo.Pwd)
	params.Add("id", userpo.ID)

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
func (dao *Dao) List(listReqVo ListReqVo) (common.PageRespVo, error) {
	args := []interface{}{}

	var whereSQL string
	if listReqVo.Active == "true" {
		whereSQL = "WHERE " + userpo.Status + "=?"
		args = append(args, userstatus.Active)
	}

	page := mysql.Page{
		TableName:   userpo.Table,
		ColumnNames: []string{userpo.ID, userpo.Acc, userpo.Name, userpo.RoleCode, userpo.Status},
		UniqueKey:   userpo.PK,
		OrderBy:     userpo.ID,
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
