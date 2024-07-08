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
	// s.storage.DeleteSession(req.AccessToken)

	// s.logger.Infof("user with session %s uccessfully logged out", req.AccessToken)

	return nil
}

func (s *GuardService) Check(ctx context.Context, req *requests.CheckRequest) (res *responses.CheckResponse, err error) {
	const op = "appGuard.Check"

	// session, err := s.storage.CheckSession(req.AccessToken)
	// if err != nil {
	// 	s.logger.WithError(err).Warnf("can't storage.CheckSession %s", op)

	// 	return mapper.CreateResponseCheck(false, 0),
	// 		fmt.Errorf("can't storage.CheckSession %s: %w", op, err)
	// }

	// s.logger.Infof("user %d is authorized", session.MemberID)

	// return mapper.CreateResponseCheck(true, session.MemberID), nil

	// return &responses.CheckResponse{Valid: true, MemberID: 0}, nil
	return nil, nil
}
