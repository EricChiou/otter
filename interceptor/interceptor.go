package interceptor

import (
	"otter/acl"
	"otter/constants/api"
	"otter/service/jwt"

	"github.com/valyala/fasthttp"
)

// Interceptor check jwt
func Interceptor(ctx *fasthttp.RequestCtx, aclCodes ...acl.Code) (jwt.Payload, bool, api.RespStatus) {
	var payload jwt.Payload
	var result bool

	auth := string(ctx.Request.Header.Peek(api.TokenHeader))
	if result = (len(auth) >= len(api.TokenPrefix)); !result {
		return payload, false, api.TokenError
	}

	if payload, result = jwt.Verify(auth[len(api.TokenPrefix):]); !result {
		return payload, false, api.TokenError
	}

	// check permission
	for _, aclCode := range aclCodes {
		if result = acl.Check(aclCode, payload.RoleCode); !result {
			return payload, false, api.PermissionDenied
		}
	}

	return payload, true, ""
}
