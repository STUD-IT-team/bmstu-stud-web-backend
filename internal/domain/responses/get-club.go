package responses

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type GetClub struct {
	ID          int              `"json:id"`
	Name        string           `"json:name"`
	ShortName   string           `"json:short_name"`
	Description string           `"json:description"`
	Type        string           `"json:"type"`
	Logo        domain.MediaFile `"json:"logo"`
	VkUrl       string           `"json:"vk_url"`
	TgUrl       string           `"json:"tg_url"`
	Orgs        []domain.ClubOrg
}
