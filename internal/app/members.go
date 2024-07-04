package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type membersServiceStorage interface {
	GetAllMembers(ctx context.Context) ([]domain.Member, error)
}

type MembersService struct {
	storage membersServiceStorage
}

func NewMembersService(storage membersServiceStorage) *MembersService {
	return &MembersService{storage: storage}
}

func (s *MembersService) GetAllMembers(ctx context.Context) (*responses.GetAllMembers, error) {
	var res []domain.Member
	var err error

	res, err = s.storage.GetAllMembers(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetAllMembers: %v", err)
	}

	return mapper.MakeResponseAllMembers(res), nil
}
