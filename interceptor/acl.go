package interceptor

import (
	"errors"
	"otter/acl"
	"otter/constants/api"
	"otter/service/jwt"

	"github.com/valyala/fasthttp"
)

// Acl interceptor
func Acl(ctx *fasthttp.RequestCtx, payload jwt.Payload, aclCodes ...acl.Code) error {
	if len(payload.RoleCode) <= 0 {
		return errors.New(string(api.PermissionDenied))
	}

	// check permission
	for _, aclCode := range aclCodes {
		if ok := acl.Check(aclCode, payload.RoleCode); !ok {
			return errors.New(string(api.PermissionDenied))
		}
	}

	return nil
}
