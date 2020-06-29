package constants

// Authorization Bearer Token
const (
	TokenHeader string = "Authorization"
	TokenPrefix string = "Bearer "
)

// ApiResult api result
type ApiResult string

const (
	RSSuccess          ApiResult = "ok"
	RSDuplicate        ApiResult = "duplicate"
	RSPermissionDenied ApiResult = "permissionDenied"
	RSTokenError       ApiResult = "tokenError"
	RSFormatError      ApiResult = "formatError"
	RSParseError       ApiResult = "parseError"
	RSDBError          ApiResult = "dbError"
	RSDataError        ApiResult = "dataError"
	RSServerError      ApiResult = "serverError"
	RSUnknownError     ApiResult = "unknownError"
)

// JwtAlgTyp jwt alg type
type JwtAlgTyp string

const (
	JWTHS256 JwtAlgTyp = "HS256"
	JWTHS384 JwtAlgTyp = "HS384"
	JWTHS512 JwtAlgTyp = "HS512"
)
