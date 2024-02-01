package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/consts"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GuardServiceStorage interface {
	GetUserID(ctx context.Context, user domain.User) (userID string, err error)
}

type GuardServiceCache[K comparable, V any] interface {
	Put(id K, value V)
	Delete(id K)
	Find(id K) *V
}

type GuardService struct {
	logger       *logrus.Logger
	storage      GuardServiceStorage
	sessionCache GuardServiceCache[string, domain.Session]
	grpc2.UnimplementedGuardServer
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func NewGuardService(log *logrus.Logger, storage GuardServiceStorage, sessionCache GuardServiceCache[string, domain.Session]) *GuardService {
	return &GuardService{
		logger:       log,
		storage:      storage,
		sessionCache: sessionCache,
	}
}

const sessionDurationHours = 5

func (s *GuardService) Login(ctx context.Context, req *requests.LoginRequest) (res *responses.LoginResponse, err error) {
	const op = "appGuard.Login"

	userID, err := s.storage.GetUserID(ctx, domain.User{
		Email:        req.Email,
		HashPasswrod: req.Password})
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			s.logger.Warn("user not found", err)

			return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		s.logger.WithError(err).Warnf("failed to get userID")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sessionID := uuid.NewString()

	session := domain.Session{
		UserID:    userID,
		ExpireAt:  time.Now().Add(time.Hour * time.Duration(sessionDurationHours)),
		EnteredAt: time.Now(),
	}

	s.sessionCache.Put(sessionID, session)

	s.logger.WithField("op", op).Infof("user %s logged in successfully", userID)

	return mapper.CreateResponseLogin(sessionID, session.ExpireAt.Format(consts.GrpcTimeFormat)), nil
}

func (s *GuardService) Logout(ctx context.Context, req *requests.LogoutRequest) error {
	const op = "appGuard.Logout"

	accessToken := req.AccessToken

	s.sessionCache.Delete(accessToken)

	s.logger.WithField("op", op).Infof("user with session %s uccessfully logged out", accessToken)

	return nil
}

func (s *GuardService) Check(ctx context.Context, req *requests.CheckRequest) (res *responses.CheckResponse, err error) {
	const op = "appGuard.Logout"

	accessToken := req.AccessToken

	session := s.sessionCache.Find(accessToken)
	if session == nil {
		s.logger.WithField("op", op).Info("session not found")

		return &responses.CheckResponse{
			Valid:  false,
			UserID: "",
		}, nil
	}

	if session.ExpireAt.Before(time.Now()) {
		s.logger.WithField("op", op).Info("session expired")

		s.sessionCache.Delete(session.UserID)

		return &responses.CheckResponse{
			Valid:  false,
			UserID: "",
		}, nil
	}

	s.logger.WithField("op", op).Infof("user %s is authorized", session.UserID)

	return &responses.CheckResponse{
		Valid:  true,
		UserID: session.UserID,
	}, nil
}
