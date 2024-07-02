package domain

type Club struct {
	ID          int    `"db:id"`
	Name        string `"db:name"`
	ShortName   string `"db:short_name"`
	Description string `"db:description"`
	Type        string `"db:"type"`
	LogoId      int    `"db:"logo"`
	VkUrl       string `"db:"vk_url"`
	TgUrl       string `"db:"tg_url"`
	Orgs        []ClubOrg
}

type ClubOrg struct {
	Member
	RoleName string `"db:"role_name"`
	RoleSpec string `"db:"role_spec"`
}
