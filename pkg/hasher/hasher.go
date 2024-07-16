package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

var ErrMismatchedHashAndPassword = bcrypt.ErrMismatchedHashAndPassword

func CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}
