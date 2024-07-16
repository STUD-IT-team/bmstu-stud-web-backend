package mapper

import (
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllMembers(f []domain.Member, membersMediaFiles map[int]domain.MediaFile) (*responses.GetAllMembers, error) {
	members := make([]responses.Member, 0, len(f))
	for _, v := range f {
		media, ok := membersMediaFiles[v.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media for member id %v", v.MediaID)
		}
		members = append(members,
			responses.Member{
				ID:       v.ID,
				Login:    v.Login,
				Media:    media,
				Telegram: v.Telegram,
				Vk:       v.Vk,
				Name:     v.Name,
				IsAdmin:  v.IsAdmin,
			})
	}

	return &responses.GetAllMembers{Members: members}, nil
}

func MakeResponseMember(f *domain.Member, membersMediaFile *domain.MediaFile) (*responses.GetMember, error) {
	return &responses.GetMember{
		ID:       f.ID,
		Login:    f.Login,
		Media:    *membersMediaFile,
		Telegram: f.Telegram,
		Vk:       f.Vk,
		Name:     f.Name,
		IsAdmin:  f.IsAdmin,
	}, nil
}

func MakeResponseMembersByName(f []domain.Member, membersMediaFiles map[int]domain.MediaFile) (*responses.GetMembersByName, error) {
	members := make([]responses.Member, 0, len(f))
	for _, v := range f {
		media, ok := membersMediaFiles[v.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media for member id %v", v.MediaID)
		}
		members = append(members,
			responses.Member{
				ID:       v.ID,
				Login:    v.Login,
				Media:    media,
				Telegram: v.Telegram,
				Vk:       v.Vk,
				Name:     v.Name,
				IsAdmin:  v.IsAdmin,
			})
	}

	return &responses.GetMembersByName{Members: members}, nil
}

func MakeRequestPostMember(f *requests.PostMember) *domain.Member {
	return &domain.Member{
		Login:    f.Login,
		MediaID:  f.MediaID,
		Telegram: f.Telegram,
		Vk:       f.Vk,
		Name:     f.Name,
		IsAdmin:  f.IsAdmin,
	}
}

func MakeRequestUpdateMember(f *requests.UpdateMember) *domain.Member {
	return &domain.Member{
		ID:       f.ID,
		Login:    f.Login,
		MediaID:  f.MediaID,
		Telegram: f.Telegram,
		Vk:       f.Vk,
		Name:     f.Name,
		IsAdmin:  f.IsAdmin,
	}
}
