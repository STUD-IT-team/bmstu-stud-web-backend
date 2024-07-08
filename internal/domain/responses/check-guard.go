package responses

type CheckResponse struct {
	MemberID int  `json:"member_id"`
	IsAdmin  bool `json:"is_admin"`
	Valid    bool `json:"valid"`
}
