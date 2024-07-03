package responses

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type GetAllClubs struct {
	Clubs []Club `json:"clubs"`
}

type Club struct {
	ID          int              `"json:id"`
	Name        string           `"json:name"`
	ShortName   string           `"json:short_name"`
	Description string           `"json:description"`
	Type        string           `"json:"type"`
	Logo        domain.MediaFile `"json:"logo"`
	VkUrl       string           `"json:"vk_url"`
	TgUrl       string           `"json:"tg_url"`
	Orgs        []ClubOrg        `"json:"org"`
}

type ClubOrg struct {
	ID   int    `"json:id"`
	Name string `"json:name"`
	Spec string `"json:spec"`
}
