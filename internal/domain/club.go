package domain

const (
	StudSovetClubID = 0
)

type Club struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortName        string `json:"short_name"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Type             string `json:"type"`
	LogoId           int    `json:"logo"`
	ParentID         int    `json:"parent_id"`
	VkUrl            string `json:"vk_url"`
	TgUrl            string `json:"tg_url"`
}

type ClubOrg struct {
	Member
	RoleName string `json:"role_name"`
	RoleSpec string `json:"role_spec"`
	ClubName string `json:"club_name"`
	ClubID   int    `json:"club_id"`
}
