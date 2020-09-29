package codemap

import (
	"otter/api/common"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/po/codemappo"
)

// Dao codemap dao
type Dao struct{}

// Add add codemap dao
func (dao *Dao) Add(addReqVo AddReqVo) error {
	run := func() interface{} {
		columnValues := map[string]interface{}{
			codemappo.Type:   addReqVo.Type,
			codemappo.Code:   addReqVo.Code,
			codemappo.Name:   addReqVo.Name,
			codemappo.SortNo: addReqVo.SortNo,
			codemappo.Enable: addReqVo.Enable,
		}

		_, err := mysql.Insert(codemappo.Table, columnValues)
		if err != nil {
			return err
		}

		return nil
	}

	err := jobqueue.Codemap.NewAddJob(run).(error)
	if err != nil {
		return err
	}

	return nil
}

// Update update codemap dao
func (dao *Dao) Update(updateReqVo UpdateReqVo) error {
	args := []interface{}{updateReqVo.Code, updateReqVo.Name, updateReqVo.Type, updateReqVo.SortNo, updateReqVo.Enable, updateReqVo.ID}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", codemappo.Table)
	params.Add("code", codemappo.Code)
	params.Add("name", codemappo.Name)
	params.Add("type", codemappo.Type)
	params.Add("sortNo", codemappo.SortNo)
	params.Add("enable", codemappo.Enable)
	params.Add("id", codemappo.ID)

	sql := "UPDATE #codemap "
	sql += "SET #code=?, #name=?, #type=?, #sortNo=?, #enable=? "
	sql += "WHERE #id=?"

	_, err := mysql.Exec(sql, params, args)
	if err != nil {
		return err
	}

	return nil
}

// Delete update codemap dao
func (dao *Dao) Delete(deleteReqVo DeleteReqVo) error {
	args := []interface{}{deleteReqVo.ID}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", codemappo.Table)
	params.Add("id", codemappo.ID)

	sql := "DELETE FROM #codemap "
	sql += "WHERE #id=?"

	_, err := mysql.Exec(sql, params, args)
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
