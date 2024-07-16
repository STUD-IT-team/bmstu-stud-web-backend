package responses

type GetAllMembers struct {
	Members []GetMember `json:"members"`
}
