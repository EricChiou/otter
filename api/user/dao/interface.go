package dao

import (
	"otter/api/user/vo"
	"otter/service/jwt"
)

// Dao user dao interface
type Dao interface {
	SignUp(signUp vo.SignUpReq) (result string, err error)
	SignIn(signIn vo.SignInReq) (signInRes vo.SignInRes, result string, err error)
	Update(payload jwt.Payload, updateData vo.UpdateReq) (result string, err error)
	List(page, limit int, where bool) (list vo.ListRes, result string, err error)
}
