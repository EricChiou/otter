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
	column := []string{
		entity.Col().ID,
		entity.Col().Type,
		entity.Col().Code,
		entity.Col().Name,
		entity.Col().SortNo,
		entity.Col().Enable,
	}
	where := map[string]interface{}{}
	if len(listReqVo.Type) > 0 {
		where[entity.Col().Type] = listReqVo.Type
	}
	if listReqVo.Enable == "true" {
		where[entity.Col().Enable] = true
	}
	orderBy := entity.Col().SortNo
	total, err := mysql.Page(entity.Table(), entity.PK(), column, where, orderBy, listReqVo.Page, listReqVo.Limit, func(result mysql.Rows) error {
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
	list.Total = total

	apihandler.Response(ctx, api.Success, list, nil)
}
