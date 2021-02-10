package auth

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("Invalid token")
	ErrTokenExpire 	= errors.New("token is expired")
)