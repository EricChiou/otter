package dao

import (
	"otter/api/user/entity"
	"otter/api/user/vo"
	cons "otter/constants"
	"otter/db/mysql"
	"otter/service/jwt"
	"otter/service/sha3"
)

// UserDao user dao
type UserDao struct {
}

// NewDao new a user dao
func NewDao() Dao {
	return &UserDao{}
}

// SignUp dao
func (dao *UserDao) SignUp(signUp vo.SignUpReq) (string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResultDBError, err
	}

	// encrypt password
	encryptPwd := sha3.Encrypt(signUp.Pwd)

	kv := map[string]interface{}{
		entity.Email: signUp.Email,
		entity.Pwd:   encryptPwd,
		entity.Name:  signUp.Name,
	}
	_, err = mysql.Insert(tx, entity.Table, kv)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}

// SignIn dao
func (dao *UserDao) SignIn(signIn vo.SignInReq) (vo.SignInRes, string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()

	var response vo.SignInRes
	var user entity.User
	column := []string{
		entity.ID,
		entity.Email,
		entity.Pwd,
		entity.Name,
		entity.Role,
	}
	where := map[string]interface{}{
		entity.Email:  signIn.Email,
		entity.Active: true,
	}
	row := mysql.QueryRow(tx, entity.Table, column, where)
	err = row.Scan(&user.ID, &user.Email, &user.Pwd, &user.Name, &user.Role)
	if err != nil {
		return response, cons.APIResultDataError, err
	}

	if user.Pwd != sha3.Encrypt(signIn.Pwd) {
		return response, cons.APIResultDataError, nil
	}

	token, _ := jwt.Generate(user.ID, user.Email, user.Name, user.Role)
	response = vo.SignInRes{
		Token: token,
	}
	return response, cons.APIResultSuccess, nil
}

// Update dao
func (dao *UserDao) Update(payload jwt.Payload, updateData vo.UpdateReq) (string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResultDBError, err
	}

	set := map[string]interface{}{}
	if len(updateData.Name) != 0 {
		set[entity.Name] = updateData.Name
	}
	if len(updateData.Pwd) != 0 {
		set[entity.Pwd] = sha3.Encrypt(updateData.Pwd)
	}
	var where map[string]interface{} = make(map[string]interface{})
	if updateData.ID != 0 {
		where[entity.ID] = updateData.ID
	} else {
		where[entity.ID] = payload.ID
	}

	_, err = mysql.Update(tx, entity.Table, set, where)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}

// List dao
func (dao *UserDao) List(page, limit int, active bool) (vo.ListRes, string, error) {
	var list vo.ListRes
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return list, cons.APIResultDBError, err
	}

	orderBy := entity.ID
	column := []string{
		entity.ID,
		entity.Email,
		entity.Name,
		entity.Role,
		entity.Active,
	}
	where := map[string]interface{}{}
	if active {
		where[entity.Active] = true
	}
	rows, err := mysql.Paging(tx, entity.Table, entity.PK, column, where, orderBy, page, limit)
	defer rows.Close()
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}

	for rows.Next() {
		var data vo.ListData
		err = rows.Scan(&data.ID, &data.Email, &data.Name, &data.Identity, &data.Active)
		if err != nil {
			return list, mysql.ErrMsgHandler(err), err
		}
		list.Records = append(list.Records, data)
	}

	var total int
	err = tx.QueryRow("SELECT COUNT(*) FROM " + entity.Table).Scan(&total)
	if err != nil {
		return list, mysql.ErrMsgHandler(err), err
	}
	list.Total = total

	return list, cons.APIResultSuccess, nil
}
