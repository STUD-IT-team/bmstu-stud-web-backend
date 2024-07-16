package responses

type GetMembersByName struct {
	Members []GetMember `json:"members"`
}
