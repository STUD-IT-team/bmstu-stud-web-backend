package domain

type Member struct {
	ID       int64  `db:"id"`
	Password []byte `db:"password"`
	Login    string `db:"login"`
	Telegram string `db:"telegram"`
	Name     string `db:"name"`
	RoleID   int    `db:"role_id"`
	IsAdmin  bool   `db:"isAdmin"`
}
