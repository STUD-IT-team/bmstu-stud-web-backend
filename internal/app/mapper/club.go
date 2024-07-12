package mapper

import (
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeMainOrg(org *domain.ClubOrg, images *map[int]domain.MediaFile) (*responses.MainOrg, error) {
	if _, ok := (*images)[org.MediaID]; !ok {
		return &responses.MainOrg{}, fmt.Errorf("can't find image for org id: %v", org.MediaID)
	}

	return &responses.MainOrg{
		ID:       org.ID,
		Name:     org.Name,
		VkUrl:    org.Vk,
		TgUrl:    org.Telegram,
		Spec:     org.RoleSpec,
		RoleName: org.RoleName,
		Image:    (*images)[org.MediaID],
	}, nil
}

func MakeSubOrg(org *domain.ClubOrg, images *map[int]domain.MediaFile) (*responses.SubClubOrg, error) {
	if _, ok := (*images)[org.MediaID]; !ok {
		return &responses.SubClubOrg{}, fmt.Errorf("can't find image for org id: %v", org.MediaID)
	}

	return &responses.SubClubOrg{
		ID:          org.Member.ID,
		Name:        org.Member.Name,
		SubClubName: org.ClubName,
		VkUrl:       org.Vk,
		TgUrl:       org.Telegram,
		Spec:        org.RoleSpec,
		Image:       (*images)[org.MediaID],
	}, nil

}

func MakeResponseClub(club *domain.Club, mainOrgs *[]domain.ClubOrg, subOrgs *[]domain.ClubOrg, images *map[int]domain.MediaFile) (*responses.GetClub, error) {
	if _, ok := (*images)[club.LogoId]; !ok {
		return &responses.GetClub{}, fmt.Errorf("can't get logo media for club id: %v", club.LogoId)
	}

	r := &responses.GetClub{
		ID:          club.ID,
		Name:        club.Name,
		ShortName:   club.ShortName,
		Logo:        (*images)[club.LogoId],
		Description: club.Description,
		Type:        club.Type,
		VkUrl:       club.VkUrl,
		TgUrl:       club.TgUrl,
	}

	for _, org := range *mainOrgs {
		m, err := MakeMainOrg(&org, images)
		if err != nil {
			return nil, err
		}
		r.MainOrgs = append(r.MainOrgs, *m)
	}

	for _, org := range *subOrgs {
		s, err := MakeSubOrg(&org, images)
		if err != nil {
			return nil, err
		}
		r.SubOrgs = append(r.SubOrgs, *s)
	}
	return r, nil
}

func MakeResponseAllClub(clubs []domain.Club, logos map[int]domain.MediaFile, orgs []domain.ClubOrg) (*responses.GetAllClubs, error) {
	r := &responses.GetAllClubs{}
	orgMap := make(map[int][]responses.ClubOrg)
	for _, org := range orgs {
		if _, ok := orgMap[org.ClubID]; !ok {
			orgMap[org.ClubID] = []responses.ClubOrg{}
		}
		o := responses.ClubOrg{}
		o.ID = org.Member.ID
		o.Name = org.Member.Name
		o.Spec = org.RoleName
		orgMap[org.ClubID] = append(orgMap[org.ClubID], o)
	}

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
				Orgs:        orgMap[club.ID],
			})
	}
	return r, nil
}

func MakeResponseClubMembers(clubID int, mainOrgs []domain.ClubOrg, subOrgs []domain.ClubOrg, images map[int]domain.MediaFile) (*responses.GetClubMembers, error) {
	r := &responses.GetClubMembers{
		ID: clubID,
	}

	for _, org := range mainOrgs {
		m, err := MakeMainOrg(&org, &images)
		if err != nil {
			return nil, err
		}
		r.MainOrgs = append(r.MainOrgs, *m)
	}

	for _, org := range subOrgs {
		s, err := MakeSubOrg(&org, &images)
		if err != nil {
			return nil, err
		}
		r.SubOrgs = append(r.SubOrgs, *s)
	}
	return r, nil
}

func ParsePostClub(req *requests.PostClub) (*domain.Club, []domain.ClubOrg, error) {
	c := &domain.Club{}

	c.Name = req.Name
	c.ShortName = req.ShortName
	c.Description = req.Description
	c.LogoId = req.LogoId
	c.Type = req.Type
	c.VkUrl = req.VkUrl
	c.TgUrl = req.TgUrl
	c.ParentID = req.ParentID
	// c.Orgs = req.Orgs

	oarr := []domain.ClubOrg{}
	for _, org := range req.Orgs {
		o := domain.ClubOrg{}
		o.ID = org.MemberID
		o.RoleName = org.RoleName
		o.RoleSpec = org.RoleSpec
		// o.ClubID = c.ID
		oarr = append(oarr, o)
	}

	return c, oarr, nil

}

func MakeResponseClubMediaFiles(clubID int, clubPhotos []domain.ClubPhoto, media map[int]domain.MediaFile) (*responses.GetClubMedia, error) {
	r := &responses.GetClubMedia{ID: clubID}
	for _, f := range clubPhotos {
		media, ok := media[f.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media file for club photo id: %v", f.MediaID)
		}
		m := responses.ClubMedia{
			RefNumber: f.RefNumber,
			MediaFile: media,
		}
		r.Media = append(r.Media, m)
	}
	return r, nil
}

func ParseUpdateClub(req *requests.UpdateClub) (*domain.Club, []domain.ClubOrg, error) {
	c := &domain.Club{}

	c.ID = req.ID
	c.Name = req.Name
	c.ShortName = req.ShortName
	c.Description = req.Description
	c.LogoId = req.LogoId
	c.Type = req.Type
	c.VkUrl = req.VkUrl
	c.TgUrl = req.TgUrl
	c.ParentID = req.ParentID

	oarr := []domain.ClubOrg{}
	for _, org := range req.Orgs {
		o := domain.ClubOrg{}
		o.ID = org.MemberID
		o.RoleName = org.RoleName
		o.RoleSpec = org.RoleSpec
		o.ClubID = c.ID
		oarr = append(oarr, o)
	}

	return c, oarr, nil
}
