package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
)

func CreateRequestLogin(req *grpc.LoginRequest) *requests.LoginRequest {
	return &requests.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func CreateResponseLogin(token, expires string) *responses.LoginResponse {
	return &responses.LoginResponse{
		Token:   token,
		Expires: expires,
	}
}

func CreateGPRCResponseLogin(res *responses.LoginResponse) *grpc.LoginResponse {
	return &grpc.LoginResponse{
		AccessToken: res.Token,
		Expires:     res.Expires,
	}
}

func CreateRequestLogout(req *grpc.LogoutRequest) *requests.LogoutRequest {
	return &requests.LogoutRequest{
		AccessToken: req.AccessToken,
	}
}

func CreateRequestCheck(req *grpc.CheckRequest) *requests.CheckRequest {
	return &requests.CheckRequest{
		AccessToken: req.AccessToken,
	}
}

func CreateResponseCheck(valid bool, userID string) *responses.CheckResponse {
	return &responses.CheckResponse{
		Valid:  valid,
		UserID: userID,
	}
}

func CreateGPRCResponseCheck(res *responses.CheckResponse) *grpc.CheckResponse {
	return &grpc.CheckResponse{
		Valid:  res.Valid,
		UserID: res.UserID,
	}
}
