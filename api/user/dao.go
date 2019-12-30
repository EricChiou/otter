package user

import (
	cons "otter/constants"
	"otter/db/mysql"
	"otter/service/jwt"
	"otter/service/sha3"
)

// Dao user dao
type Dao struct{}

// SignUp dao
func (dao *Dao) SignUp(signUp SignUpReqVo) (string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResultDBError, err
	}

	// encrypt password
	encryptPwd := sha3.Encrypt(signUp.Pwd)

	kv := map[string]interface{}{
		Email: signUp.Email,
		Pwd:   encryptPwd,
		Name:  signUp.Name,
	}
	_, err = mysql.Insert(tx, Table, kv)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}

// SignIn dao
func (dao *Dao) SignIn(signIn SignInReqVo) (SignInResVo, string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()

	var response SignInResVo
	var entity Entity
	column := []string{
		ID,
		Email,
		Pwd,
		Name,
		Role,
	}
	where := map[string]interface{}{
		Email:  signIn.Email,
		Active: true,
	}
	row := mysql.QueryRow(tx, Table, column, where)
	err = row.Scan(&entity.ID, &entity.Email, &entity.Pwd, &entity.Name, &entity.Role)
	if err != nil {
		return response, cons.APIResultDataError, err
	}

	if entity.Pwd != sha3.Encrypt(signIn.Pwd) {
		return response, cons.APIResultDataError, nil
	}

	token, _ := jwt.Generate(entity.ID, entity.Email, entity.Name, entity.Role)
	response = SignInResVo{
		Token: token,
	}
	return response, cons.APIResultSuccess, nil
}

// Update dao
func (dao *Dao) Update(payload jwt.Payload, updateData UpdateReqVo) (string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResultDBError, err
	}

	set := map[string]interface{}{}
	if len(updateData.Name) != 0 {
		set[Name] = updateData.Name
	}
	if len(updateData.Pwd) != 0 {
		set[Pwd] = sha3.Encrypt(updateData.Pwd)
	}
	var where map[string]interface{} = make(map[string]interface{})
	if updateData.ID != 0 {
		where[ID] = updateData.ID
	} else {
		where[ID] = payload.ID
	}

	_, err = mysql.Update(tx, Table, set, where)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}

// List dao
func (dao *Dao) List(page, limit int, active bool) (ListResVo, string, error) {
	var list ListResVo
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return list, cons.APIResultDBError, err
	}

	orderBy := ID
	column := []string{
		ID,
		Email,
		Name,
		Role,
		Active,
	}
	where := map[string]interface{}{}
	if active {
		where[Active] = true
	}
	rows, err := mysql.Paging(tx, Table, PK, column, where, orderBy, page, limit)
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

	return list, cons.APIResultSuccess, nil
}