package auth

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username 	string `json:"username"`
	Role		string `json:"role"`	
}

func (c *AuthClaims) GetUsername() string {
	return c.Username
}

func (c *AuthClaims) GetRole() string {
	return c.Role
}

type TokenInfo interface {
	GetUsername() 	string
	GetRole()		string
}