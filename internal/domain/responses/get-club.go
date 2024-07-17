package responses

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type GetClub struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	ShortName        string           `json:"short_name"`
	Description      string           `json:"description"`
	ShortDescription string           `json:"short_description"`
	Type             string           `json:"type"`
	Logo             domain.MediaFile `json:"logo"`
	VkUrl            string           `json:"vk_url"`
	TgUrl            string           `json:"tg_url"`
	MainOrgs         []MainOrg        `json:"main_orgs"`
	SubOrgs          []SubClubOrg     `json:"sub_orgs"`
}

type SubClubOrg struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	SubClubName string           `json:"sub_club_name"`
	VkUrl       string           `json:"vk_url"`
	TgUrl       string           `json:"tg_url"`
	Spec        string           `json:"spec"`
	Image       domain.MediaFile `json:"image"`
}

type MainOrg struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	VkUrl    string           `json:"vk_url"`
	TgUrl    string           `json:"tg_url"`
	Spec     string           `json:"spec"`
	RoleName string           `json:"role_name"`
	Image    domain.MediaFile `json:"image"`
}
