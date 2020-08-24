package codemap

import (
	"otter/acl"
	"otter/constants/api"
	"otter/interceptor"
	"otter/service/apihandler"
	"otter/service/paramhandler"

	"github.com/EricChiou/httprouter"
)

// Controller codemap controller
type Controller struct {
	dao Dao
}

// Add add new code map
func (con *Controller) Add(context *httprouter.Context) {
	ctx := context.Ctx

	// check token
	payload, err := interceptor.Token(ctx)
	if err != nil {
		apihandler.Response(ctx, api.TokenError, nil, nil)
		return
	}

	// check acl
	aclCode := []acl.Code{acl.AddCodemap}
	if err = interceptor.Acl(ctx, payload, aclCode...); err != nil {
		apihandler.Response(ctx, api.PermissionDenied, nil, nil)
		return
	}

	// check body format
	var addReqVo AddReqVo
	if err := paramhandler.Set(ctx, &addReqVo); err != nil {
		apihandler.Response(ctx, api.FormatError, nil, err)
		return
	}

	con.dao.Add(ctx, addReqVo)
}

// Update update codemap
func (con *Controller) Update(context *httprouter.Context) {
	ctx := context.Ctx

	// check token
	payload, err := interceptor.Token(ctx)
	if err != nil {
		apihandler.Response(ctx, api.TokenError, nil, nil)
		return
	}

	// check jwt and acl
	aclCode := []acl.Code{acl.UpdateCodemap}
	if err = interceptor.Acl(ctx, payload, aclCode...); err != nil {
		apihandler.Response(ctx, api.PermissionDenied, nil, nil)
		return
	}

	// check body format
	var updateReqVo UpdateReqVo
	if err := paramhandler.Set(ctx, &updateReqVo); err != nil {
		apihandler.Response(ctx, api.FormatError, nil, err)
		return
	}

	con.dao.Update(ctx, updateReqVo)
}

// Delete delete codemap
func (con *Controller) Delete(context *httprouter.Context) {
	ctx := context.Ctx

	// check token
	payload, err := interceptor.Token(ctx)
	if err != nil {
		apihandler.Response(ctx, api.TokenError, nil, nil)
		return
	}

	// check jwt and acl
	aclCode := []acl.Code{acl.DeleteCodemap}
	if err = interceptor.Acl(ctx, payload, aclCode...); err != nil {
		apihandler.Response(ctx, api.PermissionDenied, nil, nil)
		return
	}

	// check param
	var deleteReqVo DeleteReqVo
	if err := paramhandler.Set(ctx, &deleteReqVo); err != nil {
		apihandler.Response(ctx, api.FormatError, nil, err)
		return
	}

	con.dao.Delete(ctx, deleteReqVo)
}

// List get codemap list
func (con *Controller) List(context *httprouter.Context) {
	ctx := context.Ctx

	// check token
	if _, err := interceptor.Token(ctx); err != nil {
		apihandler.Response(ctx, api.TokenError, nil, nil)
		return
	}

	// check param
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
