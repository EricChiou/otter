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

var dao Dao

// SignUp user sign up controller
func SignUp(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var signUpData SignUpReq
	err := json.Unmarshal(ctx.PostBody(), &signUpData)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResultFormatError, nil, err))
		return
	}

	// check data
	result := check.Check(signUpData.Email, signUpData.Pwd, signUpData.Name)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResultFormatError, nil, nil))
		return
	}

	apiResult, err := dao.SignUp(signUpData)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, err))
}

// SignIn user sign in controller
func SignIn(context *router.Context) {
	ctx := context.Ctx

	// check body format
	signInData := SignInReq{
		Email: string(ctx.QueryArgs().Peek("email")),
		Pwd:   string(ctx.QueryArgs().Peek("pwd")),
	}

	// check data
	result := check.Check(signInData.Email, signInData.Pwd)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResultFormatError, nil, nil))
		return
	}

	response, apiResult, err := dao.SignIn(signInData)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, response, err))
}

// Update user data
func Update(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var updateData UpdateReq
	err := json.Unmarshal(ctx.PostBody(), &updateData)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResultFormatError, nil, err))
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

	// check data
	if len(updateData.Name) == 0 && len(updateData.Pwd) == 0 {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResultFormatError, nil, nil))
		return
	}

	apiResult, err := dao.Update(payload, updateData)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, err))
}

// List get user list
func List(context *router.Context) {
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
	limit, _ := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		limit = 10
	}
	active := string(ctx.QueryArgs().Peek("active"))

	list, apiResult, err := dao.List(page, limit, (active == "true"))
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, list, err))
}
