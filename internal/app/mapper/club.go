package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseClub(club *domain.Club, logo *[]byte) *responses.GetClub {
	return &responses.GetClub{
		ID:          club.ID,
		Name:        club.Name,
		ShortName:   club.ShortName,
		Logo:        *logo,
		Description: club.Description,
		Type:        club.Type,
		VkUrl:       club.VkUrl,
		TgUrl:       club.TgUrl,
		Orgs:        club.Orgs,
	}
}
