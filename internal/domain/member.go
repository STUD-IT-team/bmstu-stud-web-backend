package domain

type Member struct {
	ID           int    `json:"id"`
	HashPassword []byte `json:"hash_password"`
	Password     []byte `json:"password"`
	Login        string `json:"login"`
	MediaID      int    `json:"media_id"`
	Telegram     string `json:"telegram"`
	Vk           string `json:"vk"`
	Name         string `json:"name"`
	IsAdmin      bool   `json:"isAdmin"`
}
