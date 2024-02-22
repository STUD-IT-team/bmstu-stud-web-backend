package responses

type CheckResponse struct {
	Valid    bool  `json:"valid"`
	MemberID int64 `json:"member_id"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}
