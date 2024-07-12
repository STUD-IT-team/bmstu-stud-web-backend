package postgres

import (
	"errors"
	"fmt"
)

const uniqueConstraintViolation string = "23505"
const foreignKeyViolation string = "23503"

var (
	TextPostgresUniqueConstraintViolation = "postgres unique constraint violation"
	TextPostgresForeignKeyViolation       = "postgres foreign key violation"
	TextPostgresUnknownError              = "postgres unknown error"
)

var (
	ErrPostgresUniqueConstraintViolation = errors.New(TextPostgresUniqueConstraintViolation)
	ErrPostgresForeignKeyViolation       = errors.New(TextPostgresForeignKeyViolation)
	ErrPostgresUnknownError              = errors.New(TextPostgresUnknownError)
)

func mapPostgresError(code string) error {
	if code == uniqueConstraintViolation {
		return ErrPostgresUniqueConstraintViolation
	}
	if code == foreignKeyViolation {
		return ErrPostgresForeignKeyViolation
	}
	return ErrPostgresUnknownError
}

func wrapPostgresError(code string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %w", mapPostgresError(code), err)
}
