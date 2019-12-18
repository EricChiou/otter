package dao

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

	var user entity.User
	kv := map[string]interface{}{
		user.EmailCol(): signUp.Email,
		user.PwdCol():   encryptPwd,
		user.NameCol():  signUp.Name,
	}
	_, err = mysql.Insert(tx, user.Table(), kv)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}

// SignIn dao
func (dao *UserDao) SignIn(signIn user.SignInReq) (user.SignInRes, string, error) {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()

	var userData entity.User
	var response user.SignInRes

	// row := mysql.DB.QueryRow("SELECT id, email, pwd, name, identity FROM user WHERE email=? AND active=?", signIn.Email, true)
	var userEntity entity.User
	column := []string{
		userEntity.IDCol(),
		userEntity.EmailCol(),
		userEntity.PwdCol(),
		userEntity.NameCol(),
		userEntity.IdentityCol(),
	}
	where := map[string]interface{}{
		userEntity.EmailCol():  signIn.Email,
		userEntity.ActiveCol(): true,
	}
	row := mysql.QueryRow(tx, userEntity.Table(), column, where)
	err = row.Scan(&userData.ID, &userData.Email, &userData.Pwd, &userData.Name, &userData.Identity)
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

	var userEntity entity.User
	set := map[string]interface{}{}
	if len(updateData.Name) != 0 {
		set[userEntity.NameCol()] = updateData.Name
	}
	if len(updateData.Pwd) != 0 {
		set[userEntity.PwdCol()] = sha3.Encrypt(updateData.Pwd)
	}
	where := map[string]interface{}{
		userEntity.IDCol(): payload.ID,
	}

	_, err = mysql.Update(tx, userEntity.Table(), set, where)
	if err != nil {
		return mysql.ErrMsgHandler(err), err
	}

	return cons.APIResultSuccess, nil
}
