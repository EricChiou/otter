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

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(signUp SignUpReqVo) error {
	run := func() interface{} {
		var g mysql.Gooq

		// encrypt password
		encryptPwd := sha3.Encrypt(signUp.Pwd)

		g.SQL.Insert(userpo.Table, userpo.Acc, userpo.Pwd, userpo.Name).
			Values("?", "?", "?")
		g.AddValues(signUp.Acc, encryptPwd, signUp.Name)

		if _, err := g.Exec(); err != nil {
			return err
		}

		return nil
	}

	return jobqueue.User.NewSignUpJob(run)
}

// SignIn dao
func (dao *Dao) SignIn(signInReqVo SignInReqVo) (userbo.SignInBo, error) {
	var g mysql.Gooq
	var signInBo userbo.SignInBo

	g.SQL.Select(
		userpo.Table+"."+userpo.ID,
		userpo.Acc,
		userpo.Pwd,
		userpo.Table+"."+userpo.Name,
		userpo.RoleCode,
		userpo.Status,
		rolepo.Table+"."+rolepo.Name,
	).
		From(userpo.Table).
		Join(rolepo.Table).On(c(userpo.RoleCode).Eq(rolepo.Code)).
		Where(c(userpo.Acc).Eq("?"))
	g.AddValues(signInReqVo.Acc)

	rowMapper := func(row *sql.Row) error {
		if err := row.Scan(
			&signInBo.ID,
			&signInBo.Acc,
			&signInBo.Pwd,
			&signInBo.Name,
			&signInBo.RoleCode,
			&signInBo.Status,
			&signInBo.RoleName,
		); err != nil {
			return err
		}
		return nil
	}

	// check account existing
	if err := g.QueryRow(rowMapper); err != nil {
		return signInBo, err
	}

	return signInBo, nil
}

// Update dao
func (dao *Dao) Update(updateData UpdateReqVo) error {
	var g mysql.Gooq

	var conditions []gooq.Condition
	if len(updateData.Name) != 0 {
		conditions = append(conditions, c(userpo.Name).Eq("?"))
		g.AddValues(updateData.Name)
	}
	if len(updateData.Pwd) != 0 {
		conditions = append(conditions, c(userpo.Pwd).Eq("?"))
		g.AddValues(sha3.Encrypt(updateData.Pwd))
	}

	g.SQL.Update(userpo.Table).Set(conditions...).Where(c(userpo.ID).Eq("?"))
	g.AddValues(updateData.ID)

	if _, err := g.Exec(); err != nil {
		return err
	}

	return nil
}

// List dao
func (dao *Dao) List(listReqVo ListReqVo) (common.PageRespVo, error) {
	index := (listReqVo.Page - 1) * listReqVo.Limit

	var g mysql.Gooq
	g.SQL.
		Select(userpo.ID, userpo.Acc, userpo.Name, userpo.RoleCode, userpo.Status).
		From(userpo.Table).
		Join("").Lp().
		/**/ Select(userpo.PK).From(userpo.Table).
		/**/ OrderBy(userpo.ID).
		/**/ Limit(strconv.Itoa(index), strconv.Itoa(listReqVo.Limit)).
		Rp().As("t").
		Using(userpo.PK)

	if listReqVo.Active == "true" {
		g.SQL.Where(c(userpo.Status).Eq("?"))
		g.AddValues(string(userstatus.Active))
	}

	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
		Total:   0,
	}
	rowMapper := func(rows *sql.Rows) error {
		for rows.Next() {
			var record ListResVo
			err := rows.Scan(&record.ID, &record.Acc, &record.Name, &record.RoleCode, &record.Status)
			if err != nil {
				return err
			}
			list.Records = append(list.Records, record)
		}
		return nil
	}

	if err := g.Query(rowMapper); err != nil {
		return list, err
	}

	var countG mysql.Gooq
	countG.SQL.Select(f.Count("*")).From(userpo.Table)
	if listReqVo.Active == "true" {
		countG.SQL.Where(c(userpo.Status).Eq("?"))
		countG.AddValues(string(userstatus.Active))
	}

	countRowMapper := func(row *sql.Row) error {
		return row.Scan(&(list.Total))
	}

	if err := countG.QueryRow(countRowMapper); err != nil {
		return list, err
	}

	return list, nil
}
