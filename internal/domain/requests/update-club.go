package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateClub struct {
	ID int `json:"id"`
	PostClub
}

type UpdateClubPointer struct {
	ID *int `json:"id"`
	PostClubPointer
}

func (p *UpdateClub) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi on club_id in request: %w", err)
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := UpdateClubPointer{}

	err = decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostMember.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostMember.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*p = UpdateClub{
		ID: id,
		PostClub: PostClub{
			Name:             *pf.Name,
			ShortName:        *pf.ShortName,
			Description:      *pf.Description,
			ShortDescription: *pf.ShortDescription,
			Type:             *pf.Type,
			LogoId:           *pf.LogoId,
			VkUrl:            *pf.VkUrl,
			TgUrl:            *pf.TgUrl,
			ParentID:         *pf.ParentID,
			Orgs:             pf.Orgs,
		},
	}

	return p.validate()
}

func (f *UpdateClub) validate() error {
	return nil
}

func (pc *UpdateClubPointer) validate() error {
	if pc.Name == nil {
		return fmt.Errorf("require: name")
	}
	if pc.ShortName == nil {
		return fmt.Errorf("require: short_name")
	}
	if pc.Description == nil {
		return fmt.Errorf("require: description")
	}
	if pc.ShortDescription == nil {
		return fmt.Errorf("require: short_description")
	}
	if pc.Type == nil {
		return fmt.Errorf("require: type")
	}
	if pc.LogoId == nil {
		return fmt.Errorf("require: logo_id")
	}
	if pc.VkUrl == nil {
		return fmt.Errorf("require: vk_url")
	}
	if pc.TgUrl == nil {
		return fmt.Errorf("require: tg_url")
	}
	if pc.ParentID == nil {
		return fmt.Errorf("require: parent_id")
	}
	if pc.Orgs == nil {
		return fmt.Errorf("require: orgs")
	}
	return nil
}
