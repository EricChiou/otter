package constants

// setting
const (
	Sha3Len        int    = 256
	TokenPrefix    string = "Bearer "
	ConfigFilePath string = "./config.ini"
)

// API result
const (
	APIResultSuccess          string = "success"
	APIResultFormatError      string = "formatError"
	APIResultPermissionDenied string = "permissionDenied"
	APIResultDBError          string = "dbError"
	APIResultServerError      string = "serverError"
	APIResultDuplicate        string = "duplicate"
	APIResultDataError        string = "dataError"
	APIResultTokenError       string = "tokenError"
)

// jwt alg type
const (
	JWTHS256 string = "HS256"
	JWTHS384 string = "HS384"
	JWTHS512 string = "HS512"
)
