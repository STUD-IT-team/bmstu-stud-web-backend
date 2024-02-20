package domain

type Role struct {
	ID       int    `db:"id"`
	RoleName string `db:"role_name"`
	RoleSpec string `db:"role_spec"`
}
