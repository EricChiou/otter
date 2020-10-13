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

	"github.com/EricChiou/gooq"
)

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(signUp SignUpReqVo) error {
	run := func() interface{} {
		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)

		var sql gooq.SQL
		sql.Insert(userpo.Table, userpo.Acc, userpo.Pwd, userpo.Name).
			Values(s(signUp.Acc), s(encryptPwd), s(signUp.Name))

		if _, err := mysql.Exec(sql.GetSQL()); err != nil {
			return err
		}

		return nil
	}

	return jobqueue.User.NewSignUpJob(run)
}

// SignIn dao
func (dao *Dao) SignIn(signInReqVo SignInReqVo) (userbo.SignInBo, error) {
	var signInBo userbo.SignInBo

	var sql gooq.SQL
	sql.Select(userpo.ID, userpo.Acc, userpo.Pwd, userpo.Name, userpo.RoleCode, userpo.Status, rolepo.Name).
		From(userpo.Table).
		Join(rolepo.Table).On(c(userpo.RoleCode).Eq(rolepo.Code)).
		Where(c(userpo.Acc).Eq(s(signInReqVo.Acc)))

	err := mysql.QueryRow(sql.GetSQL(), func(result mysql.Row) error {
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
	var conditions []gooq.Condition
	if len(updateData.Name) != 0 {
		conditions = append(conditions, c(userpo.Name).Eq(s(updateData.Name)))
	}
	if len(updateData.Pwd) != 0 {
		conditions = append(conditions, c(userpo.Pwd).Eq(s(sha3.Encrypt(updateData.Pwd))))
	}

	var sql gooq.SQL
	sql.Update(userpo.Table).Set(conditions...).Where(c(userpo.ID).Eq(updateData.ID))

	_, err := mysql.Exec(sql.GetSQL())
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
