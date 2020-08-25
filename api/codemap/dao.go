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
		kvParams := mysql.GetSQLParamsInstance()
		kvParams.Add(entity.Col().Type, addReqVo.Type)
		kvParams.Add(entity.Col().Code, addReqVo.Code)
		kvParams.Add(entity.Col().Name, addReqVo.Name)
		kvParams.Add(entity.Col().SortNo, addReqVo.SortNo)
		kvParams.Add(entity.Col().Enable, addReqVo.Enable)

		_, err := mysql.Insert(entity.Table(), kvParams)
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
	setParams := mysql.GetSQLParamsInstance()
	setParams.Add(entity.Col().Code, updateReqVo.Code)
	setParams.Add(entity.Col().Name, updateReqVo.Name)
	setParams.Add(entity.Col().Type, updateReqVo.Type)
	setParams.Add(entity.Col().SortNo, updateReqVo.SortNo)
	setParams.Add(entity.Col().Enable, updateReqVo.Enable)

	whereParams := mysql.GetSQLParamsInstance()
	whereParams.Add(entity.Col().ID, updateReqVo.ID)

	_, err := mysql.Update(entity.Table(), setParams, whereParams)
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

	whereParams := mysql.GetSQLParamsInstance()
	whereParams.Add(entity.Col().ID, deleteReqVo.ID)

	_, err := mysql.Delete(entity.Table(), whereParams)
	if err != nil {
		apihandler.Response(ctx, mysql.ErrMsgHandler(err), nil, err)
		return
	}

	apihandler.Response(ctx, api.Success, nil, nil)
}

// List get codemap list
func (dao *Dao) List(ctx *fasthttp.RequestCtx, listReqVo ListReqVo) {
	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
		Total:   0,
	}

	var entity Entity
	params := mysql.GetSQLParamsInstance()
	params.Add("codemapT", entity.Table())
	params.Add("pk", entity.PK())
	params.Add("idCol", entity.Col().ID)
	params.Add("typeCol", entity.Col().Type)
	params.Add("codeCol", entity.Col().Code)
	params.Add("nameCol", entity.Col().Name)
	params.Add("sortNoCol", entity.Col().SortNo)
	params.Add("enableCol", entity.Col().Enable)
	params.Add("index", (listReqVo.Page-1)*listReqVo.Limit)
	params.Add("limit", listReqVo.Limit)

	whereParams := mysql.GetSQLParamsInstance()
	if len(listReqVo.Type) > 0 {
		whereParams.Add("#typeCol", ":type")
		params.Add("type", listReqVo.Type)
	}
	if listReqVo.Enable == "true" {
		whereParams.Add("#enableCol", ":enable")
		params.Add("enable", true)
	}
	whereSQL := mysql.WhereSQL(whereParams)

	sql := ""
	sql += "SELECT #idCol, #typeCol, #codeCol, #nameCol, #sortNoCol, #enableCol "
	sql += "FROM #codemapT "
	sql += "INNER JOIN ( "
	sql += "    SELECT #pk FROM #codemapT " + whereSQL
	sql += "    ORDER BY #idCol "
	sql += "    LIMIT :index, :limit "
	sql += ") t "
	sql += "USING ( #pk )"

	err := mysql.Query(sql, params, func(result mysql.Rows) error {
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

	sql = "SELECT COUNT(*) FROM #codemapT " + whereSQL
	var total int
	err = mysql.QueryRow(sql, params, func(result mysql.Row) error {
		return result.Row.Scan(&total)
	})
	if err != nil {
		apihandler.ResponsePage(ctx, mysql.ErrMsgHandler(err), list, err)
		return
	}
	list.Total = total

	apihandler.Response(ctx, api.Success, list, nil)
}
