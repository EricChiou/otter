package user

import (
	"fmt"

	"otter/acl"
	"otter/constants/api"
	"otter/interceptor"
	"otter/pkg/router"
	"otter/service/apihandler"
	"otter/service/paramhandler"
)

// Controller user controller
type Controller struct {
	dao Dao
}

// SignUp user sign up controller
func (con *Controller) SignUp(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var signUpData SignUpReqVo
	if err := paramhandler.Set(ctx, &signUpData); err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.FormatError, nil, err))
		return
	}

	con.dao.SignUp(ctx, signUpData)
}

// SignIn user sign in controller
func (con *Controller) SignIn(context *router.Context) {
	ctx := context.Ctx

	// set param
	var signInData SignInReqVo
	if err := paramhandler.Set(ctx, &signInData); err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.FormatError, nil, nil))
		return
	}

	con.dao.SignIn(ctx, signInData)
}

// Update user data
func (con *Controller) Update(context *router.Context) {
	ctx := context.Ctx

	// check jwt and acl
	aclCode := []acl.Code{acl.UpdateUser}
	payload, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, apihandler.Result(ctx, reason, nil, nil))
		return
	}

	// check body format
	var updateData UpdateReqVo
	if err := paramhandler.Set(ctx, &updateData); err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.FormatError, nil, err))
		return
	}

	con.dao.Update(ctx, payload, updateData)
}

// List get user list
func (con *Controller) List(context *router.Context) {
	ctx := context.Ctx

	// check jwt
	_, result, reason := interceptor.Interceptor(ctx)
	if !result {
		fmt.Fprintf(ctx, apihandler.Result(ctx, reason, nil, nil))
		return
	}

	// check body format
	var listReqVo ListReqVo
	if err := paramhandler.Set(ctx, &listReqVo); err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.FormatError, nil, err))
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
