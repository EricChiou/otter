package codemap

import (
	"otter/api/common"
	cons "otter/constants"
	"otter/db/mysql"
)

// Dao codemap dao
type Dao struct{}

// Add add codemap dao
func (dao *Dao) Add(addReqVo AddReqVo) (string, interface{}) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.RSDBError, err
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
		return mysql.ErrMsgHandler(err), err
	}

	return cons.RSSuccess, nil
}

// Update update codemap dao
func (dao *Dao) Update(updateReqVo UpdateReqVo) (string, interface{}) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.RSDBError, err
	}

	var entity Entity
	setKV := map[string]interface{}{
		entity.Col().Code:   updateReqVo.Code,
		entity.Col().Name:   updateReqVo.Name,
		entity.Col().SortNo: updateReqVo.SortNo,
		entity.Col().Enable: updateReqVo.Enable,
	}
	whereKV := map[string]interface{}{
		entity.Col().ID: updateReqVo.ID,
	}
	_, err = mysql.Update(tx, entity.Table(), setKV, whereKV)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.RSSuccess, nil
}

// Delete update codemap dao
func (dao *Dao) Delete(deleteReqVo DeleteReqVo) (string, interface{}) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.RSDBError, err
	}

	var entity Entity
	whereKV := map[string]interface{}{
		entity.Col().ID: deleteReqVo.ID,
	}
	_, err = mysql.Delete(tx, entity.Table(), whereKV)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.RSSuccess, nil
}

// List get codemap list
func (dao *Dao) List(page, limit int, typ string, enble bool) (common.PageRespVo, string, interface{}) {
	list := common.PageRespVo{
		Records: []interface{}{},
		Page:    page,
		Limit:   limit,
	}
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return list, cons.RSDBError, err
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
	if len(typ) > 0 {
		where[entity.Col().Type] = typ
	}
	if enble {
		where[entity.Col().Enable] = true
	}
	orderBy := entity.Col().SortNo
	rows, err := mysql.Page(tx, entity.Table(), entity.PK(), column, where, orderBy, page, limit)
	defer rows.Close()
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}

	for rows.Next() {
		var data ListDataVo
		err = rows.Scan(&data.ID, &data.Type, &data.Code, &data.Name, &data.SortNo, &data.Enable)
		if err != nil {
			return list, mysql.ErrMsgHandler(err), err
		}
		list.Records = append(list.Records, data)
	}

	var total int
	var args []interface{}
	whereStr, args := mysql.WhereString(where, args)
	err = tx.QueryRow("SELECT COUNT(*) FROM "+entity.Table()+whereStr, args...).Scan(&total)
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}
	list.Total = total

	return list, cons.RSSuccess, nil
}
