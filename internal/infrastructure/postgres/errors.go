package postgres

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

const (
	uniqueConstraintViolation string = "23505"
	foreignKeyViolation       string = "23503"
	notFoundError             string = "20000"
)

var (
	TextPostgresUniqueConstraintViolation = "postgres unique constraint violation"
	TextPostgresForeignKeyViolation       = "postgres foreign key violation"
	TextPostgresNotFoundError             = "postgres not found error"
	TextPostgresUnknownError              = "postgres unknown error"
)

var (
	ErrPostgresUniqueConstraintViolation = errors.New(TextPostgresUniqueConstraintViolation)
	ErrPostgresForeignKeyViolation       = errors.New(TextPostgresForeignKeyViolation)
	ErrPostgresUnknownError              = errors.New(TextPostgresUnknownError)
	ErrPostgresNotFoundError             = errors.New(TextPostgresNotFoundError)
)

func mapPostgresError(code string) error {
	if code == uniqueConstraintViolation {
		return ErrPostgresUniqueConstraintViolation
	}
	if code == foreignKeyViolation {
		return ErrPostgresForeignKeyViolation
	}
	if code == notFoundError {
		return ErrPostgresNotFoundError
	}
	return ErrPostgresUnknownError
}

func wrapPostgresError(err error) error {
	if err == nil {
		return nil
	}
	perr, ok := err.(pgx.PgError)
	if !ok {
		return err
	}
	return fmt.Errorf("%s %w: %w", perr.Code, mapPostgresError(perr.Code), err)
}
