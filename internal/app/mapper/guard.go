package mapper

import (
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

//func CreateRequestLogin(req *grpc.LoginRequest) *requests.LoginRequest {
//	//return &requests.LoginRequest{
//	//	Login:    req.Login,
//	//	Password: req.Password,
//	//}
//}

func CreateResponseLogin(token int64) *responses.LoginResponse {
	return &responses.LoginResponse{
		AccessToken: strconv.FormatInt(token, 10),
	}
}

//func CreateGPRCResponseLogin(res *responses.LoginResponse) *grpc.LoginResponse {
//	return &grpc.LoginResponse{
//		AccessToken: res.Token,
//		Expires:     res.Expires,
//	}
//}

//func CreateRequestLogout(req *grpc.LogoutRequest) *requests.LogoutRequest {
//	//return &requests.LogoutRequest{
//	//	AccessToken: req.AccessToken,
//	//}
//	return nil
//}

//func CreateRequestCheck(req *grpc.CheckRequest) *requests.CheckRequest {
//	//return &requests.CheckRequest{
//	//	AccessToken: req.AccessToken,
//	//}
//	return nil
//}

func CreateResponseCheck(valid bool, sess domain.Session) *responses.CheckResponse {
	return &responses.CheckResponse{
		Valid:    valid,
		MemberID: sess.MemberID,
		IsAdmin:  sess.IsAdmin,
	}
}

//func CreateGPRCResponseCheck(res *responses.CheckResponse) *grpc.CheckResponse {
//	//return &grpc.CheckResponse{
//	//	Valid:    res.Valid,
//	//	MemberID: res.MemberID,
//	//}
//	return nil
//}

func MapRegisterToMember(req *requests.Register) *domain.Member {
	return &domain.Member{
		Login:    req.Login,
		Password: []byte(req.Password),
		Name:     req.Name,
		Telegram: req.Telegram,
		Vk:       req.Vk,
		IsAdmin:  false,
	}
}
