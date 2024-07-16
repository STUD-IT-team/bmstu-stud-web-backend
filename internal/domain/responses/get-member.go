package responses

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type GetMember struct {
	ID       int              `json:"id"`
	Login    string           `json:"login"`
	Media    domain.MediaFile `json:"media"`
	Telegram string           `json:"telegram"`
	Vk       string           `json:"vk"`
	Name     string           `json:"name"`
	IsAdmin  bool             `json:"isAdmin"`
}
