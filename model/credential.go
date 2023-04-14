package authModel

import (
	authUtils "github.com/ReygaFitra/auth-jwt/utils"
)

var SecretKey = authUtils.DotEnv("SECRET_KEY")
var JwtKey = []byte(SecretKey)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}