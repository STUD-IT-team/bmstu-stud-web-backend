package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	cache2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/adapters/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/consts"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/storage"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GuardService struct {
	logger       *logrus.Logger
	storage      storage.GuardStorage
	sessionCache cache.ICache[string, cache2.Session]
	grpc2.UnimplementedGuardServer
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func NewGuardService(log *logrus.Logger, storage storage.GuardStorage, sessionCache cache.ICache[string, cache2.Session]) *GuardService {
	return &GuardService{
		logger:       log,
		storage:      storage,
		sessionCache: sessionCache,
	}
}

func (s *GuardService) Login(ctx context.Context, req *requests.LoginRequest) (res *responses.LoginResponse, err error) {
	const op = "appGuard.Login"

	userID, err := s.storage.GetUserID(context.TODO(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			s.logger.Warn("user not found", err)

			return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		s.logger.WithError(err).Warnf("failed to get userID")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sessionID := uuid.NewString()

	var sessionDuration int
	if sessionDuration, err = strconv.Atoi(os.Getenv("SESSION_DURATION_HOURS")); err != nil {
		sessionDuration = 5
	}

	session := cache2.Session{
		UserID:    userID,
		ExpireAt:  time.Now().Add(time.Hour * time.Duration(sessionDuration)),
		EnteredAt: time.Now(),
	}

	s.sessionCache.Put(sessionID, session)

	s.logger.WithField("op", op).Info("user logged in successfully")

	return &responses.LoginResponse{
		Token:   sessionID,
		Expires: session.ExpireAt.Format(consts.GrpcTimeFormat),
	}, nil

}

func (s *GuardService) Logout(ctx context.Context, req *requests.LogoutRequest) error {
	const op = "appGuard.Logout"

	accessToken := req.AccessToken

	s.sessionCache.Delete(accessToken)

	s.logger.WithField("op", op).Info("user successfully logged out")

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

	s.logger.WithField("op", op).Info("user is authorized")

	return &responses.CheckResponse{
		Valid:  true,
		UserID: session.UserID,
	}, nil
}
