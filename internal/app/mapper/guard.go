package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
)

func CreateRequestLogin(req *grpc2.LoginRequest) *requests.LoginRequest {
	return &requests.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func CreateGPRCResponseLogin(res *responses.LoginResponse) *grpc2.LoginResponse {
	return &grpc2.LoginResponse{
		AccessToken: res.Token,
		Expires:     res.Expires,
	}
}

func CreateRequestLogout(req *grpc2.LogoutRequest) *requests.LogoutRequest {
	return &requests.LogoutRequest{
		AccessToken: req.AccessToken,
	}
}

func CreateRequestCheck(req *grpc2.CheckRequest) *requests.CheckRequest {
	return &requests.CheckRequest{
		AccessToken: req.AccessToken,
	}
}

func CreateGPRCResponseCheck(res *responses.CheckResponse) *grpc2.CheckResponse {
	return &grpc2.CheckResponse{
		Valid:  res.Valid,
		UserID: res.UserID,
	}
}
