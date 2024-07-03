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

func MakeResponseAllClub(clubs []domain.Club, logos map[int]domain.MediaFile) (*responses.GetAllClubs, error) {
	r := &responses.GetAllClubs{}
	for _, club := range clubs {
		if _, ok := logos[club.LogoId]; !ok {
			return nil, fmt.Errorf("can't find logo for club id %v", logos)
		}
		r.Clubs = append(r.Clubs,
			responses.Club{
				ID:          club.ID,
				Name:        club.Name,
				ShortName:   club.ShortName,
				Logo:        logos[club.LogoId],
				Description: club.Description,
				Type:        club.Type,
				VkUrl:       club.VkUrl,
				TgUrl:       club.TgUrl,
				Orgs:        club.Orgs,
			})
	}
	return r, nil
}
