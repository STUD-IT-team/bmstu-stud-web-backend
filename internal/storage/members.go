package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type memberStorage interface {
	GetMemberByLogin(ctx context.Context, login string) (domain.Member, error)
	GetAllMembers(ctx context.Context) ([]domain.Member, error)
	GetMember(ctx context.Context, id int) (domain.Member, error)
	GetMembersByName(ctx context.Context, name string) ([]domain.Member, error)
	PostMember(ctx context.Context, member domain.Member) error
	DeleteMember(ctx context.Context, id int) error
	UpdateMember(ctx context.Context, member domain.Member) error
}

func (s *storage) GetMemberByLogin(ctx context.Context, login string) (domain.Member, error) {
	return s.postgres.GetMemberByLogin(ctx, login)
}

func (s *storage) GetAllMembers(ctx context.Context) ([]domain.Member, error) {
	return s.postgres.GetAllMembers(ctx)
}

func (s *storage) GetMember(ctx context.Context, id int) (domain.Member, error) {
	return s.postgres.GetMember(ctx, id)
}

func (s *storage) GetMembersByName(ctx context.Context, name string) ([]domain.Member, error) {
	return s.postgres.GetMembersByName(ctx, name)
}

func (s *storage) PostMember(ctx context.Context, member domain.Member) error {
	return s.postgres.PostMember(ctx, member)
}

func (s *storage) DeleteMember(ctx context.Context, id int) error {
	return s.postgres.DeleteMember(ctx, id)
}

func (s *storage) UpdateMember(ctx context.Context, member domain.Member) error {
	return s.postgres.UpdateMember(ctx, member)
}
