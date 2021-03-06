package jwt

import (
	"encoding/json"
	"strconv"
	"time"

	"otter/config"
	"otter/pkg/jwt"
)

// Payload jwt payload struct
type Payload struct {
	ID       int    `json:"id"`
	Acc      string `json:"acc"`
	Name     string `json:"name"`
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
	Exp      int64  `json:"exp"`
}

// Generate generate jwt
func Generate(userID int, acc, name, roleCode, roleName string) (string, error) {
	cfg := config.Get()
	jwtExpire, err := strconv.Atoi(cfg.JWTExpire)
	if err != nil {
		jwtExpire = 1
	}

	payload := Payload{
		ID:       userID,
		Acc:      acc,
		Name:     name,
		RoleCode: roleCode,
		RoleName: roleName,
		Exp:      time.Now().Unix() + int64(jwtExpire*86400),
	}
	return jwt.GenerateJWT(payload, string(config.JwtAlg), cfg.JWTKey)
}

// Verify verify JWT
func Verify(jwtStr string) (Payload, bool) {
	cfg := config.Get()
	var payload Payload
	bytes, result := jwt.VerifyJWT(jwtStr, string(config.JwtAlg), cfg.JWTKey)
	if !result {
		return payload, false
	}

	if json.Unmarshal(bytes, &payload) != nil {
		return payload, false
	}

	if time.Now().Unix() > payload.Exp {
		return payload, false
	}

	return payload, true
}
