package user

import (
	"database/sql"
	"otter/api/common"
	"otter/bo/userbo"
	"otter/constants/userstatus"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/po/rolepo"
	"otter/po/userpo"
	"otter/service/sha3"
	"strconv"

	"github.com/EricChiou/gooq"
)

type Row struct {
	Rows *sql.Rows
}

var rows *sql.Rows

// Dao user dao
type Dao struct {
	gooq mysql.Gooq
}

// SignUp dao
func (dao *Dao) SignUp(signUp SignUpReqVo) error {
	run := func() interface{} {
		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)

		var SQL gooq.SQL
		SQL.Insert(userpo.Table, userpo.Acc, userpo.Pwd, userpo.Name).
			Values(s(signUp.Acc), s(encryptPwd), s(signUp.Name))

		if _, err := dao.gooq.Exec(SQL.GetSQL(), nil); err != nil {
			return err
		}

		return nil
	}

	return jobqueue.User.NewSignUpJob(run)
}

// SignIn dao
func (dao *Dao) SignIn(signInReqVo SignInReqVo) (userbo.SignInBo, error) {
	var signInBo userbo.SignInBo

	var SQL gooq.SQL
	SQL.Select(userpo.Table+"."+userpo.ID, userpo.Acc, userpo.Pwd, userpo.Table+"."+userpo.Name, userpo.RoleCode, userpo.Status, rolepo.Table+"."+rolepo.Name).
		From(userpo.Table).
		Join(rolepo.Table).On(c(userpo.RoleCode).Eq(rolepo.Code)).
		Where(c(userpo.Acc).Eq(s(signInReqVo.Acc)))

	err := dao.gooq.QueryRow(SQL.GetSQL(), func(row *sql.Row) error {
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

	var SQL gooq.SQL
	SQL.Update(userpo.Table).Set(conditions...).Where(c(userpo.ID).Eq(updateData.ID))

	_, err := dao.gooq.Exec(SQL.GetSQL(), nil)
	if err != nil {
		return err
	}

	return nil
}

// List dao
func (dao *Dao) List(listReqVo ListReqVo) (common.PageRespVo, error) {
	index := (listReqVo.Page - 1) * listReqVo.Limit

	var SQL gooq.SQL
	var subSQL gooq.SQL
	SQL.Select(userpo.ID, userpo.Acc, userpo.Name, userpo.RoleCode, userpo.Status).
		From(userpo.Table).
		Join(subSQL.Lp().
			Select(userpo.PK).From(userpo.Table).
			OrderBy(userpo.ID).
			Limit(strconv.Itoa(index), strconv.Itoa(listReqVo.Limit)).
			Rp().GetSQL()).As("t").
		Using(userpo.PK)

	if listReqVo.Active == "true" {
		SQL.Where(c(userpo.Status).Eq(s(string(userstatus.Active))))
	}

	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
		Total:   0,
	}

	if err := dao.gooq.Query(SQL.GetSQL(), func(rows *sql.Rows) error {
		for rows.Next() {
			var record ListResVo
			err := rows.Scan(&record.ID, &record.Acc, &record.Name, &record.RoleCode, &record.Status)
			if err != nil {
				return err
			}
			list.Records = append(list.Records, record)
		}
		return nil

	}); err != nil {
		return list, err
	}

	var countSQL gooq.SQL
	countSQL.Select(f.Count("*")).From(userpo.Table)
	if listReqVo.Active == "true" {
		countSQL.Where(c(userpo.Status).Eq(s(string(userstatus.Active))))
	}

	var total int
	if err := dao.gooq.QueryRow(countSQL.GetSQL(), func(row *sql.Row) error {
		return row.Scan(&total)

	}); err != nil {
		return list, err
	}
	list.Total = total

	return list, nil
}
