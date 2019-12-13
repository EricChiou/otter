package constants

type setting struct {
	Sha3Len     int
	TokenPrefix string
}

type apiResult struct {
	Success          string
	FormatError      string
	PermissionDenied string
	DBError          string
	ServerError      string
	Duplicate        string
	DataError        string
	TokenError       string
}

type jwtAlg struct {
	HS256 string
	HS384 string
	HS512 string
}
