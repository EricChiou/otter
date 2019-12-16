package interceptor

import (
	cons "otter/constants"
	"otter/service/jwt"

	"github.com/valyala/fasthttp"
)

// Interceptor check jwt
func Interceptor(ctx *fasthttp.RequestCtx) (jwt.Payload, bool) {
	var payload jwt.Payload
	auth := string(ctx.Request.Header.Peek("Authorization"))
	if len(auth) < len(cons.Setting.TokenPrefix) {
		return payload, false
	}

	payload, result := jwt.Verify(auth[len(cons.Setting.TokenPrefix):])
	if !result {
		return payload, false
	}

	return payload, true
}
