package user

import (
	"errors"
)

// err for users
var (
	ErrUserExsist 	= errors.New("User with this username exsist")
	ErrUserNotFound	= errors.New("User not found")
)