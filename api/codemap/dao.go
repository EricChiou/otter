package codemap

import (
	"otter/api/common"
	"otter/constants/api"
	"otter/db/mysql"
	"otter/jobqueue/queues"
	"otter/service/apihandler"

	"github.com/valyala/fasthttp"
)

// Dao codemap dao
type Dao struct{}

// Add add codemap dao
func (dao *Dao) Add(ctx *fasthttp.RequestCtx, addReqVo AddReqVo) {
	wait := make(chan int)
	queues.Codemap.Add.Add(func() {
		defer func() {
			wait <- 1
		}()

		var entity Entity
		columnValues := map[string]interface{}{
			entity.Col().Type:   addReqVo.Type,
			entity.Col().Code:   addReqVo.Code,
			entity.Col().Name:   addReqVo.Name,
			entity.Col().SortNo: addReqVo.SortNo,
			entity.Col().Enable: addReqVo.Enable,
		}

		_, err := mysql.Insert(entity.Table(), columnValues)
		if err != nil {
			apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
			return
		}

		apihandler.Response(ctx, api.Success, nil, nil)
	})
	<-wait
}

// Update update codemap dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, updateReqVo UpdateReqVo) {
	var entity Entity

	args := []interface{}{updateReqVo.Code, updateReqVo.Name, updateReqVo.Type, updateReqVo.SortNo, updateReqVo.Enable, updateReqVo.ID}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", entity.Table())
	params.Add("code", entity.Col().Code)
	params.Add("name", entity.Col().Name)
	params.Add("type", entity.Col().Type)
	params.Add("sortNo", entity.Col().SortNo)
	params.Add("enable", entity.Col().Enable)
	params.Add("id", entity.Col().ID)

	sql := "UPDATE #codemap "
	sql += "SET #code=?, #name=?, #type=?, #sortNo=?, #enable=? "
	sql += "WHERE #id=?"

	_, err := mysql.Exec(sql, params, args)
	if err != nil {
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
		return
	}

	apihandler.Response(ctx, api.Success, nil, nil)
	return
}

// Delete update codemap dao
func (dao *Dao) Delete(ctx *fasthttp.RequestCtx, deleteReqVo DeleteReqVo) {
	var entity Entity

	args := []interface{}{deleteReqVo.ID}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", entity.Table())
	params.Add("id", entity.Col().ID)

	sql := "DELETE FROM #codemap "
	sql += "WHERE #id=?"

	_, err := mysql.Exec(sql, params, args)
	if err != nil {
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
		return
	}

	apihandler.Response(ctx, api.Success, nil, nil)
}

// List get codemap list
func (dao *Dao) List(ctx *fasthttp.RequestCtx, listReqVo ListReqVo) {
	var entity Entity

	args := []interface{}{(listReqVo.Page - 1) * listReqVo.Limit, listReqVo.Limit}
	whereArgs := []interface{}{}

	params := mysql.SQLParamsInstance()
	params.Add("codemap", entity.Table())
	params.Add("pk", entity.PK())
	params.Add("id", entity.Col().ID)
	params.Add("type", entity.Col().Type)
	params.Add("code", entity.Col().Code)
	params.Add("name", entity.Col().Name)
	params.Add("sortNo", entity.Col().SortNo)
	params.Add("enable", entity.Col().Enable)

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
	sql += "INNER JOIN ( "
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
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), list, err)
		return
	}

	sql = "SELECT COUNT(*) FROM #codemap " + whereSQL
	var total int
	err = mysql.QueryRow(sql, params, whereArgs, func(result mysql.Row) error {
		return result.Row.Scan(&total)
	})
	if err != nil {
		apihandler.ResponsePage(ctx, mysql.ErrMsgHandler(err), list, err)
		return
	}
	list.Total = total

	apihandler.Response(ctx, api.Success, list, nil)
}
