package codemap

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

// Add add new code map
func Add(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var addReqVo AddReqVo
	err := json.Unmarshal(ctx.PostBody(), &addReqVo)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, err))
		return
	}

	// check data
	result := check.Check(addReqVo.Type, addReqVo.Code, addReqVo.Name)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	// check jwt and acl
	aclCode := []string{acl.AddCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	apiResult, trace := dao.Add(addReqVo)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, trace))
}

// Update update codemap
func Update(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var updateReqVo UpdateReqVo
	err := json.Unmarshal(ctx.PostBody(), &updateReqVo)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, err))
		return
	}

	// check data
	result := check.Check(updateReqVo.ID, updateReqVo.Code, updateReqVo.Name)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	// check jwt and acl
	aclCode := []string{acl.UpdateCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	apiResult, trace := dao.Update(updateReqVo)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, trace))
}

// Delete delete codemap
func Delete(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var deleteReqVo DeleteReqVo
	err := json.Unmarshal(ctx.PostBody(), &deleteReqVo)
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, err))
		return
	}

	// check data
	result := check.Check(deleteReqVo.ID)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	// check jwt and acl
	aclCode := []string{acl.DeleteCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	apiResult, trace := dao.Delete(deleteReqVo)
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, nil, trace))
}

// List get codemap list
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
	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		limit = 10
	}
	typ := string(ctx.QueryArgs().Peek("type"))
	enable := string(ctx.QueryArgs().Peek("enable"))

	list, apiResult, trace := dao.List(page, limit, typ, (enable == "true"))
	fmt.Fprintf(ctx, api.Result(ctx, apiResult, list, trace))
}
