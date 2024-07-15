package domain

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrIsExpired        = errors.New("expired")
	ErrIncorrectRequest = errors.New("incorrect request")
)
