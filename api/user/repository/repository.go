package repository

import (
	"otter/api/user"
	cons "otter/constants"
	"otter/db"
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
	tx, err := db.MySQL.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResult.DBError, err
	}

	// encrypt password
	encryptPwd := sha3.Encrypt(signUp.Pwd)
	_, err = tx.Exec("INSERT INTO user( email, pwd, name ) values( ?, ?, ? )", signUp.Email, encryptPwd, signUp.Name)
	if err != nil {
		return db.MySQLErrMsgHandler(err), err
	}

	return cons.APIResult.Success, nil
}

// SignIn dao
func (dao *UserDao) SignIn(signIn user.SignInReq) (user.SignInRes, string, error) {
	var userData entity.User
	var response user.SignInRes

	row := db.MySQL.QueryRow("SELECT id, email, pwd, name, identity FROM user WHERE email=? AND active=?", signIn.Email, true)
	err := row.Scan(&userData.ID, &userData.Email, &userData.Pwd, &userData.Name, &userData.Identity)
	if err != nil {
		return response, cons.APIResult.DataError, err
	}

	if userData.Pwd != sha3.Encrypt(signIn.Pwd) {
		return response, cons.APIResult.DataError, nil
	}

	token, _ := jwt.Generate(userData.ID, userData.Email, userData.Name, userData.Identity)
	response = user.SignInRes{
		Token: token,
	}
	return response, cons.APIResult.Success, nil
}

// Update dao
func (dao *UserDao) Update(payload jwt.Payload, updateData user.UpdateReq) (string, error) {
	tx, err := db.MySQL.Begin()
	defer tx.Commit()
	if err != nil {
		return cons.APIResult.DBError, err
	}

	sql := ""
	if len(updateData.Name) != 0 {
		sql += ", name='" + updateData.Name + "'"
	}
	if len(updateData.Pwd) != 0 {
		sql += ", pwd='" + sha3.Encrypt(updateData.Pwd) + "'"
	}
	sql = sql[1:]

	_, err = tx.Exec("UPDATE user SET"+sql+" WHERE id=?", payload.ID)
	if err != nil {
		return db.MySQLErrMsgHandler(err), err
	}

	return cons.APIResult.Success, nil
}
