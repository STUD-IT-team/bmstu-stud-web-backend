package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/consts"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/hasher"
	"github.com/sirupsen/logrus"
)

type guardServiceStorage interface {
	GetUserByEmail(_ context.Context, email string) (domain.User, error)
	SetSession(id string, value domain.Session)
	FindSession(id string) *domain.Session
	DeleteSession(id string)
	SaveSessoinFromUserID(userID string) (sessionID string, session domain.Session)
}

type GuardService struct {
	logger  *logrus.Logger
	storage guardServiceStorage
	grpc.UnimplementedGuardServer
}

func NewGuardService(log *logrus.Logger, storage guardServiceStorage) *GuardService {
	return &GuardService{
		logger:  log,
		storage: storage,
	}
}

func (s *GuardService) Login(ctx context.Context, req *requests.LoginRequest) (res *responses.LoginResponse, err error) {
	const op = "appGuard.Login"

	user, err := s.storage.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			s.logger.WithError(err).Warnf("can't storage.GetUserID %s", op)
			return nil, fmt.Errorf("can't storage.GetUserID %s: %w", op, err)
		}

		s.logger.WithError(err).Warnf("can't storage.GetUserID %s", op)
		return nil, fmt.Errorf("can't storage.GetUserID %s: %w", op, err)
	}

	err = hasher.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		s.logger.WithError(err).Warnf("can't hasher.CompareHashAndPassword %s", op)
		if errors.Is(err, hasher.ErrMismatchedHashAndPassword) {
			return nil, fmt.Errorf("can't hasher.CompareHashAndPassword %s: %w", op, err)
		}

		return nil, fmt.Errorf("can't hasher.CompareHashAndPassword %s: %w", op, err)
	}

	sessionID, session := s.storage.SaveSessoinFromUserID(user.ID)

	s.logger.Infof("user %s logged in successfully", user.Email)

	return mapper.CreateResponseLogin(sessionID, session.ExpireAt.Format(consts.GrpcTimeFormat)), nil
}

func (s *GuardService) Logout(ctx context.Context, req *requests.LogoutRequest) error {

	s.storage.DeleteSession(req.AccessToken)

	s.logger.Infof("user with session %s uccessfully logged out", req.AccessToken)

	return nil
}

func (s *GuardService) Check(ctx context.Context, req *requests.CheckRequest) (res *responses.CheckResponse, err error) {
	const op = "appGuard.Logout"

	session := s.storage.FindSession(req.AccessToken)
	if session == nil {
		s.logger.WithError(domain.ErrNotFound).Warnf("can't storage.FindSession %s", op)

		return mapper.CreateResponseCheck(false, ""),
			fmt.Errorf("can't storage.FindSession %s: %w", op, domain.ErrNotFound)
	}

	if session.IsExpired() {
		s.logger.WithError(domain.ErrNotFound).Warnf("can't session.IsExpired %s", op)

		s.storage.DeleteSession(session.UserID)

		return mapper.CreateResponseCheck(false, ""),
			fmt.Errorf("can't session.IsExpired %s: %w", op, domain.ErrNotFound)
	}

	s.logger.Infof("user %s is authorized", session.UserID)

	return mapper.CreateResponseCheck(true, session.UserID), nil
}
