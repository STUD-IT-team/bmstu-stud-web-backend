package responses

type GetAllMembers struct {
	Members []Member `json:"members"`
}

type Member struct {
	ID           int    `db:"id"`
	HashPassword []byte `db:"hash_password"`
	Login        string `db:"login"`
	MediaID      int    `db:"media_id"`
	Telegram     string `db:"telegram"`
	Vk           string `db:"vk"`
	Name         string `db:"name"`
	RoleID       int    `db:"role_id"`
	IsAdmin      bool   `db:"isAdmin"`
}
