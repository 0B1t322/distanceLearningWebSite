package server

import (
	"errors"
)

var (
	ErrIncorrectUserNamePass = errors.New("Incorrect username/password")
)