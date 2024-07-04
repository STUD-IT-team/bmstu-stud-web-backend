package responses

type GetMembersByName struct {
	Members []Member `json:"members"`
}
