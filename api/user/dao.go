package user

import (
	"otter/service/jwt"
)

// Dao user dao interface
type Dao interface {
	SignUp(signUp SignUpReq) (result string, err error)
	SignIn(signIn SignInReq) (signInRes SignInRes, result string, err error)
	Update(payload jwt.Payload, updateData UpdateReq) (result string, err error)
	List(page, limit int, where bool) (list ListRes, result string, err error)
}
