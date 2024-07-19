package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostClub struct {
	Name             string    `json:"name"`
	ShortName        string    `json:"short_name"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	Type             string    `json:"type"`
	LogoId           int       `json:"logo_id"`
	VkUrl            string    `json:"vk_url"`
	TgUrl            string    `json:"tg_url"`
	ParentID         int       `json:"parent_id"`
	Orgs             []ClubOrg `json:"orgs"`
}

type PostClubPointer struct {
	Name             *string   `json:"name"`
	ShortName        *string   `json:"short_name"`
	Description      *string   `json:"description"`
	ShortDescription *string   `json:"short_description"`
	Type             *string   `json:"type"`
	LogoId           *int      `json:"logo_id"`
	VkUrl            *string   `json:"vk_url"`
	TgUrl            *string   `json:"tg_url"`
	ParentID         *int      `json:"parent_id"`
	Orgs             []ClubOrg `json:"orgs"`
}
type ClubOrg struct {
	MemberID int    `json:"member_id"`
	RoleName string `json:"role_name"`
	RoleSpec string `json:"role_spec"`
}

func (c *PostClub) Bind(req *http.Request) error {
	pc := PostClubPointer{}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&pc)

	if err != nil {
		return fmt.Errorf("can't json decoder on PostClub: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("postClub Bind: extraneous data after JSON object")
	}

	err = pc.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*c = PostClub{
		Name:             *pc.Name,
		ShortName:        *pc.ShortName,
		Description:      *pc.Description,
		ShortDescription: *pc.ShortDescription,
		Type:             *pc.Type,
		LogoId:           *pc.LogoId,
		VkUrl:            *pc.VkUrl,
		TgUrl:            *pc.TgUrl,
		ParentID:         *pc.ParentID,
		Orgs:             pc.Orgs,
	}

	return c.validate()
}

func (c *PostClub) validate() error {
	return nil
}

func (pc *PostClubPointer) validate() error {
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
