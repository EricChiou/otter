package delivery

import (
	"encoding/json"
	"fmt"

	"otter/api/user"
	"otter/api/user/repository"
	cons "otter/constants"
	"otter/interceptor"
	"otter/router"
	api "otter/service/apihandler"
	check "otter/service/checkparam"
)

var dao user.Dao = repository.NewDao()

// SignUp user sign up controller
func SignUp(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var signUpData user.SignUpReq
	err := json.Unmarshal(ctx.PostBody(), &signUpData)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResult.FormatError, nil, err))
		return
	}

	// check data
	result := check.Check(signUpData.Email, signUpData.Pwd, signUpData.Name)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResult.FormatError, nil, nil))
		return
	}

	apiResult, err := dao.SignUp(signUpData)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, err))
}

// SignIn user sign in controller
func SignIn(context *router.Context) {
	ctx := context.Ctx

	// check body format
	signInData := user.SignInReq{
		Email: string(ctx.QueryArgs().Peek("email")),
		Pwd:   string(ctx.QueryArgs().Peek("pwd")),
	}

	// check data
	result := check.Check(signInData.Email, signInData.Pwd)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResult.FormatError, nil, nil))
		return
	}

	response, apiResult, err := dao.SignIn(signInData)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, response, err))
}

// Update user data
func Update(context *router.Context) {
	ctx := context.Ctx

	// check jwt
	payload, result := interceptor.Interceptor(ctx)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResult.TokenError, nil, nil))
	}

	// check body format
	var updateData user.UpdateReq
	err := json.Unmarshal(ctx.PostBody(), &updateData)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResult.FormatError, nil, err))
		return
	}

	// check data
	if len(updateData.Name) == 0 && len(updateData.Pwd) == 0 {
		fmt.Fprintf(ctx, api.Result(ctx, cons.APIResult.FormatError, nil, nil))
		return
	}

	apiResult, err := dao.Update(payload, updateData)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, err))
}
