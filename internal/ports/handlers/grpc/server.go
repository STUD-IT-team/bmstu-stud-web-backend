package grpc

import (
	"context"
	"errors"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/hasher"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Guard interface {
	Login(ctx context.Context, req *requests.LoginRequest) (res *responses.LoginResponse, err error)
	Logout(ctx context.Context, req *requests.LogoutRequest) error
	Check(ctx context.Context, req *requests.CheckRequest) (res *responses.CheckResponse, err error)
}

type ServerAPI struct {
	grpc2.UnimplementedGuardServer
	guard Guard
}

func Register(gRPCServer *grpc.Server, guard Guard) {
	grpc2.RegisterGuardServer(gRPCServer, &ServerAPI{guard: guard})
}

func (s *ServerAPI) Login(ctx context.Context, req *grpc2.LoginRequest) (*grpc2.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	mappedReq := mapper.CreateRequestLogin(req)

	res, err := s.guard.Login(ctx, mappedReq)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) || errors.Is(err, hasher.ErrMismatchedHashAndPassword) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return mapper.CreateGPRCResponseLogin(res), nil
}

func (s *ServerAPI) Logout(ctx context.Context, req *grpc2.LogoutRequest) (*grpc2.EmptyResponse, error) {
	if err := validateLogout(req); err != nil {
		return nil, err
	}

	mappedReq := mapper.CreateRequestLogout(req)

	if err := s.guard.Logout(ctx, mappedReq); err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &grpc2.EmptyResponse{}, nil
}

func (s *ServerAPI) Check(ctx context.Context, req *grpc2.CheckRequest) (*grpc2.CheckResponse, error) {
	if err := validateCheck(req); err != nil {
		return nil, err
	}

	mappedReq := mapper.CreateRequestCheck(req)

	res, err := s.guard.Check(ctx, mappedReq)

	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return mapper.CreateGPRCResponseCheck(res), nil
}

func validateLogin(req *grpc2.LoginRequest) error {
	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

func validateLogout(req *grpc2.LogoutRequest) error {
	if req.AccessToken == "" {
		return status.Error(codes.InvalidArgument, "token is required")
	}

	return nil
}

func validateCheck(req *grpc2.CheckRequest) error {
	if req.AccessToken == "" {
		return status.Error(codes.InvalidArgument, "token is required")
	}

	return nil
}
