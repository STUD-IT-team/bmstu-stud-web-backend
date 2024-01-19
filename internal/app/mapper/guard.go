package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/request"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
)

func CreateRequestLogin(req *grpc2.LoginRequest) *request.LoginRequest {
	return &request.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func CreateGPRCResponseLogin(res *responses.LoginResponse) *grpc2.LoginResponse {
	return &grpc2.LoginResponse{
		AccessToken: res.Token,
	}
}
