package codemap

import (
	"fmt"
	"otter/api/common"
	cons "otter/constants"
	"otter/db/mysql"
	"otter/jobqueue"
	api "otter/service/apihandler"

	"github.com/valyala/fasthttp"
)

// Dao codemap dao
type Dao struct{}

// Add add codemap dao
func (dao *Dao) Add(ctx *fasthttp.RequestCtx, addReqVo AddReqVo) {
	wait := make(chan int)
	jobqueue.Codemap.Add.Add(func() {
		tx, err := mysql.DB.Begin()
		defer func() {
			tx.Commit()
			wait <- 1
		}()
		if err != nil {
			fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, nil, err))
			return
		}

		var entity Entity
		kv := map[string]interface{}{
			entity.Col().Type:   addReqVo.Type,
			entity.Col().Code:   addReqVo.Code,
			entity.Col().Name:   addReqVo.Name,
			entity.Col().SortNo: addReqVo.SortNo,
			entity.Col().Enable: addReqVo.Enable,
		}
		_, err = mysql.Insert(tx, entity.Table(), kv)
		if err != nil {
			fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
			return
		}

		fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, nil, nil))
	})
	<-wait
}

// Update update codemap dao
func (dao *Dao) Update(ctx *fasthttp.RequestCtx, updateReqVo UpdateReqVo) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, nil, err))
		return
	}

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
	_, err = mysql.Update(tx, entity.Table(), setKV, whereKV)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, nil, nil))
	return
}

// Delete update codemap dao
func (dao *Dao) Delete(ctx *fasthttp.RequestCtx, deleteReqVo DeleteReqVo) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, nil, err))
		return
	}

	var entity Entity
	whereKV := map[string]interface{}{
		entity.Col().ID: deleteReqVo.ID,
	}
	_, err = mysql.Delete(tx, entity.Table(), whereKV)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), nil, err))
		return
	}

	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, nil, nil))
}

// List get codemap list
func (dao *Dao) List(ctx *fasthttp.RequestCtx, listReqVo ListReqVo) {
	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    listReqVo.Page,
		Limit:   listReqVo.Limit,
	}
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSDBError, list, err))
		return
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
	rows, err := mysql.Page(tx, entity.Table(), entity.PK(), column, where, orderBy, listReqVo.Page, listReqVo.Limit)
	defer rows.Close()
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), list, err))
		return
	}

	for rows.Next() {
		var record ListResVo
		err = rows.Scan(&record.ID, &record.Type, &record.Code, &record.Name, &record.SortNo, &record.Enable)
		if err != nil {
			fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), list, err))
			return
		}
		list.Records = append(list.Records, record)
	}

	var total int
	var args []interface{}
	whereStr, args := mysql.WhereString(where, args)
	err = tx.QueryRow("SELECT COUNT(*) FROM "+entity.Table()+whereStr, args...).Scan(&total)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, mysql.ErrMsgHandler(err), list, err))
		return
	}
	list.Total = total

	fmt.Fprintf(ctx, api.Result(ctx, cons.RSSuccess, list, nil))
}
