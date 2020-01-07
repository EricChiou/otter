package codemap

import (
	cons "otter/constants"
	"otter/db/mysql"
	// "otter/service/jwt"
	// "otter/service/sha3"
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

	kv := map[string]interface{}{
		Col.Type:   addReqVo.Type,
		Col.Code:   addReqVo.Code,
		Col.Name:   addReqVo.Name,
		Col.SortNo: addReqVo.SortNo,
		Col.Enable: addReqVo.Enable,
	}
	_, err = mysql.Insert(tx, Table, kv)
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

	setKV := map[string]interface{}{
		Col.Code:   updateReqVo.Code,
		Col.Name:   updateReqVo.Name,
		Col.SortNo: updateReqVo.SortNo,
		Col.Enable: updateReqVo.Enable,
	}
	whereKV := map[string]interface{}{
		Col.ID: updateReqVo.ID,
	}
	_, err = mysql.Update(tx, Table, setKV, whereKV)
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

	whereKV := map[string]interface{}{
		Col.ID: deleteReqVo.ID,
	}
	_, err = mysql.Delete(tx, Table, whereKV)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.RSSuccess, nil
}

// List get codemap list
func (dao *Dao) List(page, limit int, typ string, enble bool) (ListResVo, string, interface{}) {
	var list ListResVo
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return list, cons.RSDBError, err
	}

	column := []string{
		Col.ID,
		Col.Type,
		Col.Code,
		Col.Name,
		Col.SortNo,
		Col.Enable,
	}
	where := map[string]interface{}{}
	if len(typ) > 0 {
		where[Col.Type] = typ
	}
	if enble {
		where[Col.Enable] = true
	}
	orderBy := Col.SortNo
	rows, err := mysql.Paging(tx, Table, Col.PK, column, where, orderBy, page, limit)
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
	err = tx.QueryRow("SELECT COUNT(*) FROM "+Table+whereStr, args...).Scan(&total)
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}
	list.Total = total

	return list, cons.RSSuccess, nil
}
