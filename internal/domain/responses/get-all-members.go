package responses

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type GetAllMembers struct {
	Members []Member `json:"members"`
}

type Member struct {
	ID       int              `json:"id"`
	Login    string           `json:"login"`
	Media    domain.MediaFile `json:"media"`
	Telegram string           `json:"telegram"`
	Vk       string           `json:"vk"`
	Name     string           `json:"name"`
	RoleID   int              `json:"role_id"`
	IsAdmin  bool             `json:"isAdmin"`
}
