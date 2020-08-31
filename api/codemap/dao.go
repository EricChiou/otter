package codemap

import (
	"otter/api/common"
	"otter/constants/api"
	"otter/db/mysql"
	"otter/jobqueue/queues"
	"otter/po/codemapPo"
	"otter/service/apihandler"

	"github.com/valyala/fasthttp"
)

// Dao codemap dao
type Dao struct{}

// Add add codemap dao
func (dao *Dao) Add(ctx *fasthttp.RequestCtx, addReqVo AddReqVo) apihandler.ResponseEntity {
	wait := make(chan int)
	queues.Codemap.Add.Add(func() apihandler.ResponseEntity {
		defer func() {
			wait <- 1
		}()

		columnValues := map[string]interface{}{
			codemapPo.Type:   addReqVo.Type,
			codemapPo.Code:   addReqVo.Code,
			codemapPo.Name:   addReqVo.Name,
			codemapPo.SortNo: addReqVo.SortNo,
			codemapPo.Enable: addReqVo.Enable,
		}

		_, err := mysql.Insert(codemapPo.Table, columnValues)
		if err != nil {
			return responseEntity.Error(ctx, mysql.ErrMsgHandler(err), err)
		}

		return responseEntity.OK(ctx, nil)
	})
	<-wait
	return apihandler.ResponseEntity{}
}

// Update update codemap dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, updateReqVo UpdateReqVo) apihandler.ResponseEntity {
	args := []interface{}{updateReqVo.Code, updateReqVo.Name, updateReqVo.Type, updateReqVo.SortNo, updateReqVo.Enable, updateReqVo.ID}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", codemapPo.Table)
	params.Add("code", codemapPo.Code)
	params.Add("name", codemapPo.Name)
	params.Add("type", codemapPo.Type)
	params.Add("sortNo", codemapPo.SortNo)
	params.Add("enable", codemapPo.Enable)
	params.Add("id", codemapPo.ID)

	sql := "UPDATE #codemap "
	sql += "SET #code=?, #name=?, #type=?, #sortNo=?, #enable=? "
	sql += "WHERE #id=?"

	_, err := mysql.Exec(sql, params, args)
	if err != nil {
		return responseEntity.Error(ctx, mysql.ErrMsgHandler(err), err)
	}

	return responseEntity.Error(ctx, api.Success, nil)
}

// Delete update codemap dao
func (dao *Dao) Delete(ctx *fasthttp.RequestCtx, deleteReqVo DeleteReqVo) apihandler.ResponseEntity {
	args := []interface{}{deleteReqVo.ID}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", codemapPo.Table)
	params.Add("id", codemapPo.ID)

	sql := "DELETE FROM #codemap "
	sql += "WHERE #id=?"

	_, err := mysql.Exec(sql, params, args)
	if err != nil {
		return responseEntity.Error(ctx, mysql.ErrMsgHandler(err), err)
	}

	return responseEntity.Error(ctx, api.Success, nil)
}

// List get codemap list
func (dao *Dao) List(ctx *fasthttp.RequestCtx, listReqVo ListReqVo) apihandler.ResponseEntity {
	args := []interface{}{(listReqVo.Page - 1) * listReqVo.Limit, listReqVo.Limit}
	whereArgs := []interface{}{}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", codemapPo.Table)
	params.Add("pk", codemapPo.PK)
	params.Add("id", codemapPo.ID)
	params.Add("type", codemapPo.Type)
	params.Add("code", codemapPo.Code)
	params.Add("name", codemapPo.Name)
	params.Add("sortNo", codemapPo.SortNo)
	params.Add("enable", codemapPo.Enable)

	var whereSQL string
	if len(listReqVo.Type) > 0 {
		whereSQL += "AND #type=? "
		whereArgs = append(whereArgs, listReqVo.Type)
	}
	if listReqVo.Enable == "true" {
		whereSQL += "AND #enable=? "
		whereArgs = append(whereArgs, true)
	}
	if len(whereSQL) > 0 {
		whereSQL = "WHERE " + whereSQL[4:]
	}

	sql := "SELECT #id, #type, #code, #name, #sortNo, #enable "
	sql += "FROM #codemap "
	sql += "    JOIN ( "
	sql += "    SELECT #pk FROM #codemap "
	sql += "    ORDER BY #id "
	sql += "    LIMIT ?, ? "
	sql += ") t "
	sql += "USING ( #pk )"
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
			err := rows.Scan(&record.ID, &record.Type, &record.Code, &record.Name, &record.SortNo, &record.Enable)
			if err != nil {
				return err
			}
			list.Records = append(list.Records, record)
		}

		return nil
	})
	if err != nil {
		return responseEntity.Error(ctx, mysql.ErrMsgHandler(err), err)
	}

	sql = "SELECT COUNT(*) FROM #codemap " + whereSQL
	var total int
	err = mysql.QueryRow(sql, params, whereArgs, func(result mysql.Row) error {
		return result.Row.Scan(&total)
	})
	if err != nil {
		return responseEntity.Page(ctx, mysql.ErrMsgHandler(err), list, err)
	}
	list.Total = total

	return responseEntity.OK(ctx, list)
}
