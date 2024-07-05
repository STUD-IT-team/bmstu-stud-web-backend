package responses

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type GetAllMembers struct {
	Members []Member `json:"members"`
}

type Member struct {
	ID           int              `db:"id"`
	HashPassword []byte           `db:"hash_password"`
	Login        string           `db:"login"`
	Media        domain.MediaFile `db:"media"`
	Telegram     string           `db:"telegram"`
	Vk           string           `db:"vk"`
	Name         string           `db:"name"`
	RoleID       int              `db:"role_id"`
	IsAdmin      bool             `db:"isAdmin"`
}
