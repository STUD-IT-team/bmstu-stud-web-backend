package requests

type CheckRequest struct {
	AccessToken string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LogoutRequest struct {
	AccessToken string
}
