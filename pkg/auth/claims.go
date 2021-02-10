package auth

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	TokenInfo	
}

type TokenInfo struct {
	UID			string `json:"uid"`
	Username	string `json:"username"`
	Role 		string `json:"role"`
}