package jwt

// AlgTyp jwt alg type
type AlgTyp string

const (
	HS256 AlgTyp = "HS256"
	HS384 AlgTyp = "HS384"
	HS512 AlgTyp = "HS512"
)
