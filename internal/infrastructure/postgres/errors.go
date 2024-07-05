package postgres

import (
	"errors"
	"fmt"
)

const uniqueConstraintViolation int = 23505

var (
	TextPostgresUniqueConstraintViolation = "postgres unique constraint violation"
	TextPostgresUnknownError              = "postgres unknown error"
)

var (
	ErrPostgresUniqueConstraintViolation = errors.New(TextPostgresUniqueConstraintViolation)
	ErrPostgresUnknownError              = errors.New(TextPostgresUnknownError)
)

func mapPostgresError(code int) error {
	if code == uniqueConstraintViolation {
		return ErrPostgresUniqueConstraintViolation
	}
	return ErrPostgresUnknownError
}

func wrapPostgresError(code int, err error) error {
	return fmt.Errorf("%v: %v", mapPostgresError(code), err)
}
