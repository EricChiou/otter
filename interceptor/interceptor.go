package interceptor

import (
	"otter/acl"
	cons "otter/constants"
	"otter/service/jwt"

	"github.com/valyala/fasthttp"
)

// Interceptor check jwt
func Interceptor(ctx *fasthttp.RequestCtx, aclCode ...string) (jwt.Payload, bool, string) {
	var payload jwt.Payload
	auth := string(ctx.Request.Header.Peek(cons.TokenHeader))
	if len(auth) < len(cons.TokenPrefix) {
		return payload, false, cons.APIResultTokenError
	}

	payload, result := jwt.Verify(auth[len(cons.TokenPrefix):])
	if !result {
		return payload, false, cons.APIResultTokenError
	}

	// check permission
	for _, code := range aclCode {
		return payload, acl.Check(code, payload.Role), cons.APIResultPermissionDenied
	}

	return payload, true, ""
}
