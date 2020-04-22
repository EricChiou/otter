package user

import (
	"encoding/json"
	"fmt"
	"strconv"

	"otter/acl"
	cons "otter/constants"
	"otter/interceptor"
	"otter/router"
	api "otter/service/apihandler"
	check "otter/service/checkparam"
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
	err := json.Unmarshal(ctx.PostBody(), &signUpData)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, err))
		return
	}

	// check data
	result := check.Check(signUpData.Acc, signUpData.Pwd, signUpData.Name)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	con.dao.SignUp(ctx, signUpData)
}

// SignIn user sign in controller
func (con *Controller) SignIn(context *router.Context) {
	ctx := context.Ctx

	// check body format
	signInData := SignInReqVo{
		Acc: string(ctx.QueryArgs().Peek("acc")),
		Pwd: string(ctx.QueryArgs().Peek("pwd")),
	}

	// check data
	result := check.Check(signInData.Acc, signInData.Pwd)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	con.dao.SignIn(ctx, signInData)
}

// Update user data
func (con *Controller) Update(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var updateData UpdateReqVo
	err := json.Unmarshal(ctx.PostBody(), &updateData)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, err))
		return
	}

	// check data
	result := check.Check(updateData.Name, updateData.Pwd)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	// check jwt and acl
	var aclCode []string
	if updateData.ID != 0 {
		aclCode = append(aclCode, acl.UpdateUserInfo)
	}
	payload, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	apiResult, trace := con.dao.Update(ctx, payload, updateData)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, trace))
}

// List get user list
func (con *Controller) List(context *router.Context) {
	ctx := context.Ctx

	// check jwt
	_, result, reason := interceptor.Interceptor(ctx)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	// check body format
	page, err := strconv.Atoi(string(ctx.QueryArgs().Peek("page")))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		limit = 10
	}
	active := string(ctx.QueryArgs().Peek("active"))

	list, apiResult, trace := con.dao.List(ctx, page, limit, (active == "true"))
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, list, trace))
}
