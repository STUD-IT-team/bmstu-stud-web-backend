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
	GetMember(ctx context.Context, id int) (*domain.Member, error)
	GetMembersByName(ctx context.Context, name string) ([]domain.Member, error)
	PostMember(ctx context.Context, member *domain.Member) error
	DeleteMember(ctx context.Context, id int) error
	UpdateMember(ctx context.Context, member *domain.Member) error
	GetMediaFile(ctx context.Context, id int) (*domain.MediaFile, error)
	GetMediaFiles(ctx context.Context, ids []int) (map[int]domain.MediaFile, error)
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
		return nil, fmt.Errorf("can't storage.GetAllMembers: %w", err)
	}

	ids := make([]int, 0, len(res))
	for _, member := range res {
		ids = append(ids, member.MediaID)
	}

	membersMediaFiles, err := s.storage.GetMediaFiles(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetmemberMediaFiles: %w", err)
	}

	return mapper.MakeResponseAllMembers(res, membersMediaFiles)
}

func (s *MembersService) GetMember(ctx context.Context, id int) (*responses.GetMember, error) {
	res, err := s.storage.GetMember(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetMember: %w", err)
	}

	feedMediaFile, err := s.storage.GetMediaFile(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFeedMediaFile: %w", err)
	}

	return mapper.MakeResponseMember(res, feedMediaFile)
}

func (s *MembersService) GetMembersByName(ctx context.Context, name string) (*responses.GetMembersByName, error) {
	res, err := s.storage.GetMembersByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetMembersByName: %w", err)
	}

	ids := make([]int, 0, len(res))
	for _, member := range res {
		ids = append(ids, member.MediaID)
	}

	membersMediaFiles, err := s.storage.GetMediaFiles(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetmemberMediaFiles: %w", err)
	}

	return mapper.MakeResponseMembersByName(res, membersMediaFiles)
}

func (s *MembersService) PostMember(ctx context.Context, member *domain.Member) error {
	err := s.storage.PostMember(ctx, member)
	if err != nil {
		return fmt.Errorf("can't storage.PostMember: %w", err)
	}

	return nil
}

func (s *MembersService) DeleteMember(ctx context.Context, id int) error {
	err := s.storage.DeleteMember(ctx, id)
	if err != nil {
		return fmt.Errorf("can't storage.DeleteMember: %w", err)
	}

	return nil
}

func (s *MembersService) UpdateMember(ctx context.Context, member *domain.Member) error {
	err := s.storage.UpdateMember(ctx, member)
	if err != nil {
		return fmt.Errorf("can't storage.UpdateMember: %w", err)
	}

	return nil
}

func (s *MembersService) GetClearance(ctx context.Context, resp *responses.CheckResponse) (*responses.GetClearance, error) {
	if resp.IsAdmin {
		return &responses.GetClearance{Access: true, Comment: ""}, nil
	}
	return &responses.GetClearance{Access: false, Comment: "only admins"}, nil
}
