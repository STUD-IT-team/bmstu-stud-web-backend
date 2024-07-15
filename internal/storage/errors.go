package storage

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrIsExpired         = errors.New("expired")
	ErrCantCreateSession = errors.New("cant create session")
)
