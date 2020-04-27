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
	aclCode := []acl.Code{acl.AddCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	dao.Add(ctx, addReqVo)
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
	result := check.Check(updateReqVo.ID)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	// check jwt and acl
	aclCode := []acl.Code{acl.UpdateCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	dao.Update(ctx, updateReqVo)
}

// Delete delete codemap
func Delete(context *router.Context) {
	ctx := context.Ctx

	// check param
	var deleteReqVo DeleteReqVo
	id, err := strconv.Atoi(string(ctx.QueryArgs().Peek("id")))
	if err != nil {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, err))
		return
	}
	deleteReqVo.ID = id

	// check data
	result := check.Check(deleteReqVo.ID)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, cons.RSFormatError, nil, nil))
		return
	}

	// check jwt and acl
	aclCode := []acl.Code{acl.DeleteCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, api.Result(ctx, reason, nil, nil))
		return
	}

	dao.Delete(ctx, deleteReqVo)
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

	// check param
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

	dao.List(ctx, page, limit, typ, (enable == "true"))
}
