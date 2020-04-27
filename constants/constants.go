package constants

// Authorization Bearer Token
const (
	TokenHeader string = "Authorization"
	TokenPrefix string = "Bearer "
)

// ApiResult api result
type ApiResult string

const (
	RSSuccess          ApiResult = "success"
	RSFormatError      ApiResult = "formatError"
	RSPermissionDenied ApiResult = "permissionDenied"
	RSDBError          ApiResult = "dbError"
	RSServerError      ApiResult = "serverError"
	RSDuplicate        ApiResult = "duplicate"
	RSDataError        ApiResult = "dataError"
	RSTokenError       ApiResult = "tokenError"
)

// JwtAlgTyp jwt alg type
type JwtAlgTyp string

const (
	JWTHS256 JwtAlgTyp = "HS256"
	JWTHS384 JwtAlgTyp = "HS384"
	JWTHS512 JwtAlgTyp = "HS512"
)
