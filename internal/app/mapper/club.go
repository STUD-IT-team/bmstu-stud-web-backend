package mapper

import (
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseClub(club *domain.Club, logo *domain.MediaFile) *responses.GetClub {
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

func MakeResponseAllClub(clubs []domain.Club, logos []domain.MediaFile) (*responses.GetAllClubs, error) {
	if len(clubs) != len(logos) {
		return nil, fmt.Errorf("length of club list must equal logos: %v != %v", len(clubs), len(logos))
	}
	r := &responses.GetAllClubs{}
	for i, club := range clubs {
		r.Clubs = append(r.Clubs,
			responses.Club{
				ID:          club.ID,
				Name:        club.Name,
				ShortName:   club.ShortName,
				Logo:        logos[i],
				Description: club.Description,
				Type:        club.Type,
				VkUrl:       club.VkUrl,
				TgUrl:       club.TgUrl,
				Orgs:        club.Orgs,
			})
	}
	return r, nil
}
