package responses

type GetAllEventMemberRoles struct {
	EventMemberRole []EventMemberRole `json:"roles"`
}

type EventMemberRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
