package user

import (
	"otter/api/common"
	cons "otter/constants"
	"otter/db/mysql"
	"otter/service/jwt"
	"otter/service/sha3"
)

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(signUp SignUpReqVo) (string, interface{}) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.RSDBError, err
	}

	// encrypt password
	encryptPwd := sha3.Encrypt(signUp.Pwd)

	kv := map[string]interface{}{
		Col.Email: signUp.Email,
		Col.Pwd:   encryptPwd,
		Col.Name:  signUp.Name,
	}
	_, err = mysql.Insert(tx, Table, kv)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.RSSuccess, nil
}

// SignIn dao
func (dao *Dao) SignIn(signIn SignInReqVo) (SignInResVo, string, interface{}) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()

	var response SignInResVo
	var entity Entity
	column := []string{
		Col.ID,
		Col.Email,
		Col.Pwd,
		Col.Name,
		Col.Role,
	}
	where := map[string]interface{}{
		Col.Email:  signIn.Email,
		Col.Active: true,
	}
	row := mysql.QueryRow(tx, Table, column, where)
	err = row.Scan(&entity.ID, &entity.Email, &entity.Pwd, &entity.Name, &entity.Role)
	if err != nil {
		return response, cons.RSDataError, err
	}

	if entity.Pwd != sha3.Encrypt(signIn.Pwd) {
		return response, cons.RSDataError, nil
	}

	token, _ := jwt.Generate(entity.ID, entity.Email, entity.Name, entity.Role)
	response = SignInResVo{
		Token: token,
	}
	return response, cons.RSSuccess, nil
}

// Update dao
func (dao *Dao) Update(payload jwt.Payload, updateData UpdateReqVo) (string, interface{}) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.RSDBError, err
	}

	set := map[string]interface{}{}
	if len(updateData.Name) != 0 {
		set[Col.Name] = updateData.Name
	}
	if len(updateData.Pwd) != 0 {
		set[Col.Pwd] = sha3.Encrypt(updateData.Pwd)
	}
	var where map[string]interface{} = make(map[string]interface{})
	if updateData.ID != 0 {
		where[Col.ID] = updateData.ID
	} else {
		where[Col.ID] = payload.ID
	}

	_, err = mysql.Update(tx, Table, set, where)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.RSSuccess, nil
}

// List dao
func (dao *Dao) List(page, limit int, active bool) (common.PageRespVo, string, interface{}) {
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

	column := []string{
		Col.ID,
		Col.Email,
		Col.Name,
		Col.Role,
		Col.Active,
	}
	where := map[string]interface{}{}
	if active {
		where[Col.Active] = true
	}
	orderBy := Col.ID
	rows, err := mysql.Page(tx, Table, Col.PK, column, where, orderBy, page, limit)
	defer rows.Close()
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}

	for rows.Next() {
		var data ListDataVo
		err = rows.Scan(&data.ID, &data.Email, &data.Name, &data.Identity, &data.Active)
		if err != nil {
			return list, mysql.ErrMsgHandler(err), err
		}
		list.Records = append(list.Records, data)
	}

	var total int
	err = tx.QueryRow("SELECT COUNT(*) FROM " + Table).Scan(&total)
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}
	list.Total = total

	return list, cons.RSSuccess, nil
}
