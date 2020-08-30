package interceptor

import (
	"otter/acl"
	"otter/service/apihandler"
	"otter/service/jwt"

	"github.com/EricChiou/httprouter"
)

type WebInput struct {
	ctx     *httprouter.Context
	payload jwt.Payload
}

func interceptor(ctx *httprouter.Context, needToken bool, aclCodes acl.Code, run func(WebInput) apihandler.ResponseEntity) apihandler.ResponseEntity {
	// check token

	// check acl

	webInput := WebInput{
		ctx: ctx,
	}

	return run(webInput)
}

func Get(path string, needToken bool, aclCodes acl.Code, run func(WebInput) apihandler.ResponseEntity) {
	httprouter.Get(path, func(ctx *httprouter.Context) {
		interceptor(ctx, needToken, aclCodes, run)
	})
}
