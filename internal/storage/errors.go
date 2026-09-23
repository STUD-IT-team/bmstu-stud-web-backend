package storage

import "errors"

var (
	ErrIsExpired         = errors.New("expired")
	ErrCantCreateSession = errors.New("cant create session")
)
