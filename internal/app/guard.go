package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/consts"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"github.com/sirupsen/logrus"
)

type guardServiceStorage interface {
	DeleteSession(id string)
	SaveSessoinFromUserID(userID string) (session domain.Session)
	GetUserAndValidatePassword(ctx context.Context, email string, password string) (domain.User, error)
	CheckSession(accessToken string) (*domain.Session, error)
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

	user, err := s.storage.GetUserAndValidatePassword(ctx, req.Email, req.Password)
	if err != nil {
		s.logger.WithError(err).Warnf("can't storage.GetUserAndValidatePassword %s", op)
		return nil, fmt.Errorf("can't storage.GetUserAndValidatePassword %s: %w", op, err)
	}

	session := s.storage.SaveSessoinFromUserID(user.ID)

	s.logger.Infof("user %s logged in successfully", user.Email)

	return mapper.CreateResponseLogin(session.SessionID, session.ExpireAt.Format(consts.GrpcTimeFormat)), nil
}

func (s *GuardService) Logout(ctx context.Context, req *requests.LogoutRequest) error {
	s.storage.DeleteSession(req.AccessToken)

	s.logger.Infof("user with session %s uccessfully logged out", req.AccessToken)

	return nil
}

func (s *GuardService) Check(ctx context.Context, req *requests.CheckRequest) (res *responses.CheckResponse, err error) {
	const op = "appGuard.Check"

	session, err := s.storage.CheckSession(req.AccessToken)
	if err != nil {
		s.logger.WithError(err).Warnf("can't storage.CheckSession %s", op)

		return mapper.CreateResponseCheck(false, ""),
			fmt.Errorf("can't storage.CheckSession %s: %w", op, err)
	}

	s.logger.Infof("user %s is authorized", session.UserID)

	return mapper.CreateResponseCheck(true, session.UserID), nil
}
