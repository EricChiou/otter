package interceptor

import (
	"otter/acl"
	cons "otter/constants"
	"otter/service/jwt"

	"github.com/valyala/fasthttp"
)

// Interceptor check jwt
func Interceptor(ctx *fasthttp.RequestCtx, aclCodes ...acl.Code) (jwt.Payload, bool, cons.ApiResult) {
	var payload jwt.Payload
	var result bool

	auth := string(ctx.Request.Header.Peek(cons.TokenHeader))
	if result = (len(auth) >= len(cons.TokenPrefix)); !result {
		return payload, false, cons.RSTokenError
	}

	if payload, result = jwt.Verify(auth[len(cons.TokenPrefix):]); !result {
		return payload, false, cons.RSTokenError
	}

	// check permission
	for _, aclCode := range aclCodes {
		if result = acl.Check(aclCode, payload.RoleCode); !result {
			return payload, false, cons.RSPermissionDenied
		}
	}

	return payload, true, ""
}
