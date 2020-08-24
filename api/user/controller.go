package user

import (
	"otter/acl"
	"otter/constants/api"
	"otter/interceptor"
	"otter/service/apihandler"
	"otter/service/paramhandler"

	"github.com/EricChiou/httprouter"
)

// Controller user controller
type Controller struct {
	dao Dao
}

// SignUp user sign up controller
func (con *Controller) SignUp(context *httprouter.Context) {
	ctx := context.Ctx

	// check body format
	var signUpData SignUpReqVo
	if err := paramhandler.Set(ctx, &signUpData); err != nil {
		apihandler.Response(ctx, api.FormatError, nil, err)
		return
	}

	con.dao.SignUp(ctx, signUpData)
}

// SignIn user sign in controller
func (con *Controller) SignIn(context *httprouter.Context) {
	ctx := context.Ctx

	// set param
	var signInData SignInReqVo
	if err := paramhandler.Set(ctx, &signInData); err != nil {
		apihandler.Response(ctx, api.FormatError, nil, nil)
		return
	}

	con.dao.SignIn(ctx, signInData)
}

// Update user data
func (con *Controller) Update(context *httprouter.Context) {
	ctx := context.Ctx

	// check token
	payload, err := interceptor.Token(ctx)
	if err != nil {
		apihandler.Response(ctx, api.TokenError, nil, nil)
		return
	}

	// check body format
	var updateData UpdateReqVo
	if err := paramhandler.Set(ctx, &updateData); err != nil {
		apihandler.Response(ctx, api.FormatError, nil, err)
		return
	}

	// check acl
	if updateData.ID != 0 && updateData.ID != payload.ID {
		aclCode := []acl.Code{acl.UpdateUser}
		if err := interceptor.Acl(ctx, payload, aclCode...); err != nil {
			apihandler.Response(ctx, api.PermissionDenied, nil, nil)
			return
		}
	}

	con.dao.Update(ctx, payload, updateData)
}

// List get user list
func (con *Controller) List(context *httprouter.Context) {
	ctx := context.Ctx

	// check token
	_, err := interceptor.Token(ctx)
	if err != nil {
		apihandler.Response(ctx, api.TokenError, nil, nil)
		return
	}

	// check body format
	var listReqVo ListReqVo
	if err := paramhandler.Set(ctx, &listReqVo); err != nil {
		apihandler.Response(ctx, api.FormatError, nil, err)
		return
	}

	if listReqVo.Page == 0 {
		listReqVo.Page = 1
	}
	if listReqVo.Limit == 0 {
		listReqVo.Limit = 10
	}

	con.dao.List(ctx, listReqVo)
}
