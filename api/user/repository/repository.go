package repository

import (
	"otter/api/user"
	cons "otter/constants"
	"otter/db/mysql"
	"otter/entity"
	"otter/service/jwt"
	"otter/service/sha3"
)

// UserDao user dao
type UserDao struct {
}

// NewDao new a user dao
func NewDao() user.Dao {
	return &UserDao{}
}

// SignUp dao
func (dao *UserDao) SignUp(signUp user.SignUpReq) (string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResultDBError, err
	}

	// encrypt password
	encryptPwd := sha3.Encrypt(signUp.Pwd)
	_, err = tx.Exec("INSERT INTO user( email, pwd, name ) values( ?, ?, ? )", signUp.Email, encryptPwd, signUp.Name)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}

// SignIn dao
func (dao *UserDao) SignIn(signIn user.SignInReq) (user.SignInRes, string, error) {
	var userData entity.User
	var response user.SignInRes

	row := mysql.DB.QueryRow("SELECT id, email, pwd, name, identity FROM user WHERE email=? AND active=?", signIn.Email, true)
	err := row.Scan(&userData.ID, &userData.Email, &userData.Pwd, &userData.Name, &userData.Identity)
	if err != nil {
		return response, cons.APIResultDataError, err
	}

	if userData.Pwd != sha3.Encrypt(signIn.Pwd) {
		return response, cons.APIResultDataError, nil
	}

	token, _ := jwt.Generate(userData.ID, userData.Email, userData.Name, userData.Identity)
	response = user.SignInRes{
		Token: token,
	}
	return response, cons.APIResultSuccess, nil
}

// Update dao
func (dao *UserDao) Update(payload jwt.Payload, updateData user.UpdateReq) (string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResultDBError, err
	}

	sql := ""
	var args []interface{}
	if len(updateData.Name) != 0 {
		sql += ", name=?"
		args = append(args, updateData.Name)
	}
	if len(updateData.Pwd) != 0 {
		sql += ", pwd=?"
		args = append(args, sha3.Encrypt(updateData.Pwd))
	}
	sql = sql[1:]
	args = append(args, payload.ID)

	_, err = tx.Exec("UPDATE user SET"+sql+" WHERE id=?", args...)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}
