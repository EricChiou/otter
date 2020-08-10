package codemap

import (
	"fmt"
	"otter/api/common"
	"otter/constants/api"
	"otter/db/mysql"
	"otter/jobqueue"
	"otter/service/apihandler"

	"github.com/valyala/fasthttp"
)

// Dao codemap dao
type Dao struct{}

// Add add codemap dao
func (dao *Dao) Add(ctx *fasthttp.RequestCtx, addReqVo AddReqVo) {
	wait := make(chan int)
	jobqueue.Codemap.Add.Add(func() {
		defer func() {
			wait <- 1
		}()

		var entity Entity
		kv := map[string]interface{}{
			entity.Col().Type:   addReqVo.Type,
			entity.Col().Code:   addReqVo.Code,
			entity.Col().Name:   addReqVo.Name,
			entity.Col().SortNo: addReqVo.SortNo,
			entity.Col().Enable: addReqVo.Enable,
		}
		_, err := mysql.Insert(entity.Table(), kv)
		if err != nil {
			fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
			return
		}

		fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, nil, nil))
	})
	<-wait
}

// Update update codemap dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, updateReqVo UpdateReqVo) {
	var entity Entity
	setKV := map[string]interface{}{
		entity.Col().Code:   updateReqVo.Code,
		entity.Col().Name:   updateReqVo.Name,
		entity.Col().Type:   updateReqVo.Type,
		entity.Col().SortNo: updateReqVo.SortNo,
		entity.Col().Enable: updateReqVo.Enable,
	}
	whereKV := map[string]interface{}{
		entity.Col().ID: updateReqVo.ID,
	}
	_, err := mysql.Update(entity.Table(), setKV, whereKV)
	if err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, nil, nil))
	return
}

// Delete update codemap dao
func (dao *Dao) Delete(ctx *fasthttp.RequestCtx, deleteReqVo DeleteReqVo) {
	var entity Entity
	whereKV := map[string]interface{}{
		entity.Col().ID: deleteReqVo.ID,
	}
	_, err := mysql.Delete(entity.Table(), whereKV)
	if err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, nil, nil))
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
	total, err := mysql.Page(entity.Table(), entity.PK(), column, where, orderBy, listReqVo.Page, listReqVo.Limit, func(result mysql.RowsResult) error {
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
		fmt.Fprintf(ctx, apihandler.Result(ctx, mysql.ErrMsgHandler(err), list, err))
		return
	}
	list.Total = total

	fmt.Fprintf(ctx, apihandler.Result(ctx, api.Success, list, nil))
}
