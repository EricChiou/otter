package constants

// Setting setting
var Setting = setting{
	Sha3Len:     256,
	TokenPrefix: "Bearer ",
}

// APIResult api result type
var APIResult = apiResult{
	Success:          "success",
	FormatError:      "formatError",
	PermissionDenied: "permissionDenied",
	DBError:          "dbError",
	ServerError:      "serverError",
	Duplicate:        "duplicate",
	DataError:        "dataError",
	TokenError:       "tokenError",
}

// JWTAlg jwt alg type
var JWTAlg = jwtAlg{
	HS256: "HS256",
	HS384: "HS384",
	HS512: "HS512",
}
