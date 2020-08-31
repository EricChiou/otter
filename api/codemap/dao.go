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
	args := []interface{}{}

	var whereSQL string
	if len(listReqVo.Type) > 0 {
		whereSQL += "AND " + codemapPo.Type + "=? "
		args = append(args, listReqVo.Type)
	}
	if listReqVo.Enable == "true" {
		whereSQL += "AND " + codemapPo.Enable + "=? "
		args = append(args, true)
	}
	if len(whereSQL) > 0 {
		whereSQL = "WHERE " + whereSQL[4:]
	}

	page := mysql.Page{
		TableName:   codemapPo.Table,
		ColumnNames: []string{codemapPo.ID, codemapPo.Type, codemapPo.Code, codemapPo.Name, codemapPo.SortNo, codemapPo.Enable},
		UniqueKey:   codemapPo.PK,
		OrderBy:     codemapPo.ID,
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
		return responseEntity.Page(ctx, mysql.ErrMsgHandler(err), list, err)
	}

	return responseEntity.Page(ctx, api.Success, list, nil)
}
