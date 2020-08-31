package codemap

import (
	"otter/constants/api"
	"otter/interceptor"
	"otter/service/apihandler"
	"otter/service/paramhandler"
)

// Controller codemap controller
type Controller struct {
	dao Dao
}

// Add add new code map
func (con *Controller) Add(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	// check body format
	var addReqVo AddReqVo
	if err := paramhandler.Set(webInput.Context, &addReqVo); err != nil {
		return responseEntity.Error(ctx, api.FormatError, err)
	}

	return con.dao.Add(ctx, addReqVo)
}

// Update update codemap
func (con *Controller) Update(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	// check body format
	var updateReqVo UpdateReqVo
	if err := paramhandler.Set(webInput.Context, &updateReqVo); err != nil {
		return responseEntity.Error(ctx, api.FormatError, err)
	}

	return con.dao.Update(ctx, updateReqVo)
}

// Delete delete codemap
func (con *Controller) Delete(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	// check param
	var deleteReqVo DeleteReqVo
	if err := paramhandler.Set(webInput.Context, &deleteReqVo); err != nil {
		return responseEntity.Error(ctx, api.FormatError, err)
	}

	return con.dao.Delete(ctx, deleteReqVo)
}

// List get codemap list
func (con *Controller) List(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	// check param
	var listReqVo ListReqVo
	if err := paramhandler.Set(webInput.Context, &listReqVo); err != nil {
		return responseEntity.Error(ctx, api.FormatError, err)
	}

	if listReqVo.Page == 0 {
		listReqVo.Page = 1
	}
	if listReqVo.Limit == 0 {
		listReqVo.Limit = 10
	}

	return con.dao.List(ctx, listReqVo)
}
