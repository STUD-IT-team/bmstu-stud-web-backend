package postgres

import (
	"errors"
	"fmt"
)

const uniqueConstraintViolation string = "23505"

var (
	TextPostgresUniqueConstraintViolation = "postgres unique constraint violation"
	TextPostgresUnknownError              = "postgres unknown error"
)

var (
	ErrPostgresUniqueConstraintViolation = errors.New(TextPostgresUniqueConstraintViolation)
	ErrPostgresUnknownError              = errors.New(TextPostgresUnknownError)
)

func mapPostgresError(code string) error {
	if code == uniqueConstraintViolation {
		return ErrPostgresUniqueConstraintViolation
	}
	return ErrPostgresUnknownError
}

func wrapPostgresError(code string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %w", mapPostgresError(code), err)
}
