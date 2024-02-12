package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/hasher"
)

func (s *storage) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.postgres.GetUserByEmail(ctx, email)
}

func (s *storage) GetUserAndValidatePassword(ctx context.Context, email string, password string) (domain.User, error) {
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	err = hasher.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
