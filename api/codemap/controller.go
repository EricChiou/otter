package codemap

import (
	"fmt"

	"otter/acl"
	"otter/constants/api"
	"otter/interceptor"
	"otter/pkg/router"
	"otter/service/apihandler"
	"otter/service/paramhandler"
)

// Controller codemap controller
type Controller struct {
	dao Dao
}

// Add add new code map
func (con *Controller) Add(context *router.Context) {
	ctx := context.Ctx

	// check body format
	var addReqVo AddReqVo
	if err := paramhandler.Set(ctx, &addReqVo); err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.FormatError, nil, err))
		return
	}

	// check jwt and acl
	aclCode := []acl.Code{acl.AddCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, apihandler.Result(ctx, reason, nil, nil))
		return
	}

	con.dao.Add(ctx, addReqVo)
}

// Update update codemap
func (con *Controller) Update(context *router.Context) {
	ctx := context.Ctx

	// check jwt and acl
	aclCode := []acl.Code{acl.UpdateCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, apihandler.Result(ctx, reason, nil, nil))
		return
	}

	// check body format
	var updateReqVo UpdateReqVo
	if err := paramhandler.Set(ctx, &updateReqVo); err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.FormatError, nil, err))
		return
	}

	con.dao.Update(ctx, updateReqVo)
}

// Delete delete codemap
func (con *Controller) Delete(context *router.Context) {
	ctx := context.Ctx

	// check jwt and acl
	aclCode := []acl.Code{acl.DeleteCodemap}
	_, result, reason := interceptor.Interceptor(ctx, aclCode...)
	if !result {
		fmt.Fprintf(ctx, apihandler.Result(ctx, reason, nil, nil))
		return
	}

	// check param
	var deleteReqVo DeleteReqVo
	if err := paramhandler.Set(ctx, &deleteReqVo); err != nil {
		fmt.Fprintf(ctx, apihandler.Result(ctx, api.FormatError, nil, err))
		return
	}

	con.dao.Delete(ctx, deleteReqVo)
}

// List get codemap list
func (con *Controller) List(context *router.Context) {
	ctx := context.Ctx

	// check jwt
	_, result, reason := interceptor.Interceptor(ctx)
	if !result {
		fmt.Fprintf(ctx, apihandler.Result(ctx, reason, nil, nil))
		return
	}

	// check param
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
