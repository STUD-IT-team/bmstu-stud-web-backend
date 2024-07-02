package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

//func CreateRequestLogin(req *grpc.LoginRequest) *requests.LoginRequest {
//	//return &requests.LoginRequest{
//	//	Login:    req.Login,
//	//	Password: req.Password,
//	//}
//}

func CreateResponseLogin(token, expires string) *responses.LoginResponse {
	return &responses.LoginResponse{
		Token:   token,
		Expires: expires,
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

func CreateResponseCheck(valid bool, memberID int64) *responses.CheckResponse {
	//return &responses.CheckResponse{
	//	Valid:    valid,
	//	MemberID: memberID,
	//}
	return nil
}

//func CreateGPRCResponseCheck(res *responses.CheckResponse) *grpc.CheckResponse {
//	//return &grpc.CheckResponse{
//	//	Valid:    res.Valid,
//	//	MemberID: res.MemberID,
//	//}
//	return nil
//}
