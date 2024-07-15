package domain

type Member struct {
	ID           int    `json:"id"`
	HashPassword []byte `json:"hash_password"`
	Login        string `json:"login"`
	MediaID      int    `json:"media_id"`
	Telegram     string `json:"telegram"`
	Vk           string `json:"vk"`
	Name         string `json:"name"`
	RoleID       int    `json:"role_id"`
	IsAdmin      bool   `json:"isAdmin"`
}
