package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type memberStorage interface {
	GetMemberByLogin(ctx context.Context, login string) (domain.Member, error)
}

func (s *storage) GetMemberByLogin(ctx context.Context, login string) (domain.Member, error) {
	return s.postgres.GetMemberByLogin(ctx, login)
}
