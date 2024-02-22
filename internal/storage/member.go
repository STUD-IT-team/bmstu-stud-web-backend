package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/hasher"
)

func (s *storage) GetMemberByLogin(ctx context.Context, login string) (domain.Member, error) {
	return s.postgres.GetMemberByLogin(ctx, login)
}

func (s *storage) GetMemberAndValidatePassword(ctx context.Context, login string, password string,
) (domain.Member, error) {
	user, err := s.GetMemberByLogin(ctx, login)
	if err != nil {
		return domain.Member{}, err
	}

	err = hasher.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return domain.Member{}, err
	}

	return user, nil
}
