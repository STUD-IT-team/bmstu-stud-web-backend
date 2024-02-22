package requests

type CheckRequest struct {
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}
