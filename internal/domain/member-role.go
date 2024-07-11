package domain

type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
	RoleSpec string `json:"role_spec"`
}
