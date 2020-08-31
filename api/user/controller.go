package user

import (
	"errors"
	"otter/constants/api"
	"otter/interceptor"
	"otter/service/apihandler"
	"otter/service/paramhandler"
	"strconv"
)

// Controller user controller
type Controller struct {
	dao Dao
}

// SignUp user sign up controller
func (con *Controller) SignUp(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	// check body format
	var signUpData SignUpReqVo
	if err := paramhandler.Set(webInput.Context, &signUpData); err != nil {
		return responseEntity.Error(ctx, api.FormatError, err)
	}

	return con.dao.SignUp(ctx, signUpData)
}

// SignIn user sign in controller
func (con *Controller) SignIn(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	// set param
	var signInData SignInReqVo
	if err := paramhandler.Set(webInput.Context, &signInData); err != nil {
		return responseEntity.Error(ctx, api.FormatError, nil)
	}

	return con.dao.SignIn(ctx, signInData)
}

// Update user data, POST: /user
func (con *Controller) Update(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx
	payload := webInput.Payload

	// check body format
	var updateData UpdateReqVo
	if err := paramhandler.Set(webInput.Context, &updateData); err != nil {
		return responseEntity.Error(ctx, api.FormatError, err)
	}
	if len(updateData.Name) == 0 && len(updateData.Pwd) == 0 {
		return responseEntity.Error(ctx, api.FormatError, errors.New("need name or pwd"))
	}
	updateData.ID = payload.ID

	return con.dao.Update(ctx, updateData)
}

// UpdateByUserID POST: /user/:userID
func (con *Controller) UpdateByUserID(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	var updateData UpdateReqVo

	// check body format
	if err := paramhandler.Set(webInput.Context, &updateData); err != nil {
		return responseEntity.Error(ctx, api.FormatError, err)
	}
	if len(updateData.Name) == 0 && len(updateData.Pwd) == 0 {
		return responseEntity.Error(ctx, api.FormatError, errors.New("need name or pwd"))
	}

	// check path param
	userID, err := strconv.ParseInt(webInput.Context.PathParam("userID"), 10, 64)
	if err != nil {
		return responseEntity.Error(ctx, api.FormatError, errors.New("need userID"))
	}
	updateData.ID = int(userID)

	return con.dao.Update(ctx, updateData)
}

// List get user list
func (con *Controller) List(webInput interceptor.WebInput) apihandler.ResponseEntity {
	ctx := webInput.Context.Ctx

	// check body format
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
