package codemap

import (
	"otter/api/common"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/po/codemappo"

	"github.com/EricChiou/gooq"
)

// Dao codemap dao
type Dao struct{}

// Add add codemap dao
func (dao *Dao) Add(addReqVo AddReqVo) error {
	run := func() interface{} {
		var sql gooq.SQL
		sql.Insert(codemappo.Table, codemappo.Type, codemappo.Code, codemappo.Name, codemappo.SortNo, codemappo.Enable).
			Values(s(addReqVo.Type), s(addReqVo.Code), s(addReqVo.Name), addReqVo.SortNo, addReqVo.Enable)

		if _, err := mysql.Exec(sql.GetSQL()); err != nil {
			return err
		}

		return nil
	}

	return jobqueue.Codemap.NewAddJob(run)
}

// Update update codemap dao
func (dao *Dao) Update(updateReqVo UpdateReqVo) error {
	var sql gooq.SQL
	sql.Update(codemappo.Table).
		Set(c(codemappo.Code).Eq(s(updateReqVo.Code)), c(codemappo.Name).Eq(s(updateReqVo.Name)), c(codemappo.Type).Eq(s(updateReqVo.Type)), c(codemappo.SortNo).Eq(updateReqVo.SortNo), c(codemappo.Enable).Eq(updateReqVo.Enable)).
		Where(c(codemappo.ID).Eq(updateReqVo.ID))

	_, err := mysql.Exec(sql.GetSQL())
	if err != nil {
		return err
	}

	return nil
}

// Delete update codemap dao
func (dao *Dao) Delete(deleteReqVo DeleteReqVo) error {
	var sql gooq.SQL
	sql.Delete(codemappo.Table).Where(c(codemappo.ID).Eq(deleteReqVo.ID))

	_, err := mysql.Exec(sql.GetSQL())
	if err != nil {
		return err
	}

	return nil
}

// List get codemap list
func (dao *Dao) List(listReqVo ListReqVo) (common.PageRespVo, error) {
	args := []interface{}{}

	var whereSQL string
	if len(listReqVo.Type) > 0 {
		whereSQL += "AND " + codemappo.Type + "=? "
		args = append(args, listReqVo.Type)
	}
	if listReqVo.Enable == "true" {
		whereSQL += "AND " + codemappo.Enable + "=? "
		args = append(args, true)
	}
	if len(whereSQL) > 0 {
		whereSQL = "WHERE " + whereSQL[4:]
	}

	page := mysql.Page{
		TableName:   codemappo.Table,
		ColumnNames: []string{codemappo.ID, codemappo.Type, codemappo.Code, codemappo.Name, codemappo.SortNo, codemappo.Enable},
		UniqueKey:   codemappo.PK,
		OrderBy:     codemappo.ID,
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
			err := rows.Scan(&record.ID, &record.Type, &record.Code, &record.Name, &record.SortNo, &record.Enable)
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
