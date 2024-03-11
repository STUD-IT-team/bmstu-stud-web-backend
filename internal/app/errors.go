package app

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidArgument    = errors.New("invalid argument")
)
