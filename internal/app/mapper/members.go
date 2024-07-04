package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllMembers(f []domain.Member) *responses.GetAllMembers {
	members := make([]responses.Member, 0, len(f))
	for _, v := range f {
		members = append(members,
			responses.Member{
				ID:           v.ID,
				HashPassword: v.HashPassword,
				Login:        v.Login,
				MediaID:      v.MediaID,
				Telegram:     v.Telegram,
				Vk:           v.Vk,
				Name:         v.Name,
				RoleID:       v.RoleID,
				IsAdmin:      v.IsAdmin,
			})
	}

	return &responses.GetAllMembers{Members: members}
}
