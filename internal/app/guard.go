package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type guardServiceStorage interface {
	DeleteSession(id int64)
	CreateSession(memberID int, isAdmin bool) (domain.Session, error)
	GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error)
	CheckSession(accessToken int64) (domain.Session, error)
	RegisterMember(ctx context.Context, member *domain.Member) (int, error)
	GetRandomDefaultMedia(ctx context.Context) (*domain.DefaultMedia, error)
}

type GuardService struct {
	storage guardServiceStorage
}

func NewGuardService(storage guardServiceStorage) *GuardService {
	return &GuardService{
		storage: storage,
	}
}

func (s *GuardService) Login(ctx context.Context, req *requests.LoginRequest,
) (res *responses.LoginResponse, err error) {
	const op = "appGuard.Login"

	member, err := s.storage.GetMemberAndValidatePassword(ctx, req.Login, req.Password)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetUserAndValidatePassword %s: %w", op, err)
	}

	session, err := s.storage.CreateSession(member.ID, member.IsAdmin)
	if err != nil {
		return nil, fmt.Errorf("can't storage.CreateSession %s: %w", op, err)
	}

	return mapper.CreateResponseLogin(session.SessionID), nil
}

func (s *GuardService) Logout(_ context.Context, req *requests.LogoutRequest) error {
	s.storage.DeleteSession(req.AccessToken)

	return nil
}

func (s *GuardService) Check(ctx context.Context, req *requests.CheckRequest) (res *responses.CheckResponse, err error) {
	const op = "appGuard.Check"

	session, err := s.storage.CheckSession(req.AccessToken)
	if err != nil {
		return mapper.CreateResponseCheck(false, session),
			fmt.Errorf("can't storage.CheckSession %s: %w", op, err)
	}

	return mapper.CreateResponseCheck(true, session), nil
}

func (s *GuardService) Register(ctx context.Context, member *domain.Member) error {
	defaultMedia, err := s.storage.GetRandomDefaultMedia(ctx)
	if err != nil {
		return fmt.Errorf("can't storage.GetRandomDefaultMedia: %w", err)
	}
	member.MediaID = defaultMedia.MediaID

	_, err = s.storage.RegisterMember(ctx, member)
	if err != nil {
		return fmt.Errorf("can't storage.RegisterMember: %w", err)
	}
	return err
}
