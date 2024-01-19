package app

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/repository"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Guard struct {
	logger     *logrus.Logger
	repository repository.IUserAuthStorage
	grpc2.UnimplementedGuardServer
}

func NewGuard(log *logrus.Logger, rep repository.IUserAuthStorage) *Guard {
	return &Guard{
		logger:     log,
		repository: rep,
	}
}

func (s *Guard) Login(ctx context.Context, req *grpc2.LoginRequest) (*grpc2.LoginResponse, error) {
	const op = "Guard.Login"

	s.logger.WithFields(logrus.Fields{
		"op": op,
	})

	if err := validateLogin(req); err != nil {
		return nil, err
	}

	mappedReq := mapper.CreateRequestLogin(req)

	exists, err := s.repository.CheckUser(mappedReq.Email)
	if err != nil {

	}
}

func (s *Guard) Logout(ctx context.Context, req *grpc2.LogoutRequest) (*grpc2.EmptyResponse, error) {
	panic("implement me")
}

func (s *Guard) Check(ctx context.Context, req *grpc2.CheckRequest) (*grpc2.CheckResponse, error) {
	panic("implement me")
}

func validateLogin(req *grpc2.LoginRequest) error {
	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
}
