package codemap

import (
	"database/sql"
	"otter/api/common"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/po/codemappo"
	"strconv"

	"github.com/EricChiou/gooq"
)

// Dao codemap dao
type Dao struct {
	gooq mysql.Gooq
}

// Add add codemap dao
func (dao *Dao) Add(addReqVo AddReqVo) error {
	run := func() interface{} {
		var SQL gooq.SQL
		SQL.Insert(codemappo.Table, codemappo.Type, codemappo.Code, codemappo.Name, codemappo.SortNo, codemappo.Enable).
			Values(s(addReqVo.Type), s(addReqVo.Code), s(addReqVo.Name), addReqVo.SortNo, addReqVo.Enable)

		if _, err := dao.gooq.Exec(SQL.GetSQL(), nil); err != nil {
			return err
		}

		return nil
	}

	return jobqueue.Codemap.NewAddJob(run)
}

// Update update codemap dao
func (dao *Dao) Update(updateReqVo UpdateReqVo) error {
	var SQL gooq.SQL
	SQL.Update(codemappo.Table).
		Set(c(codemappo.Code).Eq(s(updateReqVo.Code)), c(codemappo.Name).Eq(s(updateReqVo.Name)), c(codemappo.Type).Eq(s(updateReqVo.Type)), c(codemappo.SortNo).Eq(updateReqVo.SortNo), c(codemappo.Enable).Eq(updateReqVo.Enable)).
		Where(c(codemappo.ID).Eq(updateReqVo.ID))

	_, err := dao.gooq.Exec(SQL.GetSQL(), nil)
	if err != nil {
		return err
	}

	return nil
}

// Delete update codemap dao
func (dao *Dao) Delete(deleteReqVo DeleteReqVo) error {
	var SQL gooq.SQL
	SQL.Delete(codemappo.Table).Where(c(codemappo.ID).Eq(deleteReqVo.ID))

	_, err := dao.gooq.Exec(SQL.GetSQL(), nil)
	if err != nil {
		return err
	}

	return nil
}

// List get codemap list
func (dao *Dao) List(listReqVo ListReqVo) (common.PageRespVo, error) {
	index := (listReqVo.Page - 1) * listReqVo.Limit

	var SQL gooq.SQL
	var subSQL gooq.SQL
	SQL.Select(codemappo.ID, codemappo.Type, codemappo.Code, codemappo.Name, codemappo.SortNo, codemappo.Enable).
		From(codemappo.Table).
		Join(subSQL.Lp().
			Select(codemappo.PK).From(codemappo.Table).
			OrderBy(codemappo.ID).
			Limit(strconv.Itoa(index), strconv.Itoa(listReqVo.Limit)).
			Rp().GetSQL()).As("t").
		Using(codemappo.PK)

	var whereSQL gooq.SQL
	if len(listReqVo.Type) > 0 {
		whereSQL.Where(c(codemappo.Type).Eq(s(listReqVo.Type)))
	}
	if listReqVo.Enable == "true" {
		if len(whereSQL.GetSQL()) == 0 {
			SQL.Where(c(codemappo.Enable).Eq(true))
		} else {
			SQL.And(c(codemappo.Enable).Eq(true))
		}
	}
	SQL.Add(whereSQL.GetSQL())

	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
		Total:   0,
	}
	if err := dao.gooq.Query(SQL.GetSQL(), func(rows *sql.Rows) error {
		for rows.Next() {
			var record ListResVo
			err := rows.Scan(&record.ID, &record.Type, &record.Code, &record.Name, &record.SortNo, &record.Enable)
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
	countSQL.Select(f.Count("*")).From(codemappo.Table)
	if listReqVo.Enable == "true" {
		countSQL.Where(c(codemappo.Enable).Eq(true))
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
