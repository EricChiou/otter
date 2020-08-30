package interceptor

import (
	"otter/acl"
	"otter/constants/api"
	"otter/service/apihandler"
	"otter/service/jwt"

	"github.com/EricChiou/httprouter"
)

type WebInput struct {
	Context *httprouter.Context
	Payload jwt.Payload
}

func Set(context *httprouter.Context, needToken bool, aclCodes []acl.Code, run func(WebInput) apihandler.ResponseEntity) apihandler.ResponseEntity {

	// check token
	payload, err := Token(context.Ctx)
	if needToken && err != nil {
		var responseEntity apihandler.ResponseEntity
		return responseEntity.Error(context.Ctx, api.TokenError, nil)
	}

	// check acl
	if err = Acl(context.Ctx, payload, aclCodes...); err != nil {
		var responseEntity apihandler.ResponseEntity
		return responseEntity.Error(context.Ctx, api.PermissionDenied, nil)
	}

	webInput := WebInput{
		Context: context,
	}

	return run(webInput)
}
