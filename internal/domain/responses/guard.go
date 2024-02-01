package responses

type CheckResponse struct {
	Valid  bool
	UserID string
}

type LoginResponse struct {
	Token   string
	Expires string
}
