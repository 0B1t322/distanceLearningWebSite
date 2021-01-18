package auth

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	UID			string `json:"id"`
	Username 	string `json:"username"`
	Role		string `json:"role"`	
}

func (c *AuthClaims) GetID() string {
	return c.UID
}

func (c *AuthClaims) GetUsername() string {
	return c.Username
}

func (c *AuthClaims) GetRole() string {
	return c.Role
}

type TokenInfo interface {
	GetID()			string
	GetUsername() 	string
	GetRole()		string
}