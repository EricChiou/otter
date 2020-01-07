package constants

// setting
const (
	ServerName     string = "Calico & MooMoo"
	Sha3Len        int    = 256
	TokenHeader    string = "Authorization"
	TokenPrefix    string = "Bearer "
	ConfigFilePath string = "./config.ini"
)

// API result
const (
	RSSuccess          string = "success"
	RSFormatError      string = "formatError"
	RSPermissionDenied string = "permissionDenied"
	RSDBError          string = "dbError"
	RSServerError      string = "serverError"
	RSDuplicate        string = "duplicate"
	RSDataError        string = "dataError"
	RSTokenError       string = "tokenError"
)

// jwt alg type
const (
	JWTHS256 string = "HS256"
	JWTHS384 string = "HS384"
	JWTHS512 string = "HS512"
)
