package interceptor

import (
	"errors"
	"otter/constants/api"
	"otter/service/jwt"

	"github.com/valyala/fasthttp"
)

// Token interceptor
func Token(ctx *fasthttp.RequestCtx) (jwt.Payload, error) {
	auth := string(ctx.Request.Header.Peek(api.TokenHeader))
	if len(auth) < len(api.TokenPrefix) {
		return jwt.Payload{}, errors.New("token error")
	}

	if payload, ok := jwt.Verify(auth[len(api.TokenPrefix):]); !ok {
		return payload, errors.New("token error")
	} else {
		return payload, nil
	}
}
