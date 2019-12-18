package jwt

import (
	"encoding/json"
	"strconv"
	"time"

	"otter/config"
	cons "otter/constants"
	"otter/pkg/jwt"
)

// Payload jwt payload struct
type Payload struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Identity string `json:"identity"`
	Exp      int64  `json:"exp"`
}

// Generate generate jwt
func Generate(userID int, email, name, identity string) (string, error) {
	jwtExpire, err := strconv.Atoi(config.Config.JWTExpire)
	if err != nil {
		jwtExpire = 1
	}

	payload := Payload{
		ID:       userID,
		Email:    email,
		Name:     name,
		Identity: identity,
		Exp:      time.Now().Unix() + int64(jwtExpire*86400),
	}
	return jwt.GenerateJWT(payload, cons.JWTHS256, config.Config.JWTKey)
}

// Verify verify JWT
func Verify(jwtStr string) (Payload, bool) {
	var payload Payload
	bytes, result := jwt.VerifyJWT(jwtStr, cons.JWTHS256, config.Config.JWTKey)
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
