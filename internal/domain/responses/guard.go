package responses

type CheckResponse struct {
	Valid  bool   `json:"valid"`
	UserID string `json:"user_id"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}
