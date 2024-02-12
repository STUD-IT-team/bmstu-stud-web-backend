package requests

type CheckRequest struct {
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}
