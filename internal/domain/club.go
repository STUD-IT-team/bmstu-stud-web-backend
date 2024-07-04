package domain

type Club struct {
	ID          int    `"db:id"`
	Name        string `"db:"name"`
	ShortName   string `"db:"short_name"`
	Description string `"db:"description"`
	Type        string `"db:"type"`
	LogoId      int    `"db:"logo"`
	ParentID    int    `"db:"parent_id"`
	VkUrl       string `"db:"vk_url"`
	TgUrl       string `"db:"tg_url"`
}

type ClubOrg struct {
	Member
	RoleName string `"db:"role_name"`
	RoleSpec string `"db:"role_spec"`
	ClubName string `"db:"club_name"`
	ClubID   int    `"db:"club_id"`
}
