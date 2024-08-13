package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostMember struct {
	Login    string `json:"login"`
	MediaID  int    `json:"media_id"`
	Telegram string `json:"telegram"`
	Vk       string `json:"vk"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin"`
}

type PostMemberPointer struct {
	Login    *string `json:"login"`
	MediaID  *int    `json:"media_id"`
	Telegram *string `json:"telegram"`
	Vk       *string `json:"vk"`
	Name     *string `json:"name"`
	IsAdmin  *bool   `json:"is_admin"`
}

func (f *PostMember) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := PostMemberPointer{}

	err := decoder.Decode(&pf)
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

	*f = PostMember{
		Login:    *pf.Login,
		MediaID:  *pf.MediaID,
		Telegram: *pf.Telegram,
		Vk:       *pf.Vk,
		Name:     *pf.Name,
		IsAdmin:  *pf.IsAdmin,
	}

	return f.validate()
}

func (f *PostMember) validate() error {
	return nil
}

func (pf *PostMemberPointer) validate() error {
	if pf.Login == nil {
		return fmt.Errorf("require: Login")
	}
	if pf.MediaID == nil {
		return fmt.Errorf("require: MediaID")
	}
	if pf.Telegram == nil {
		return fmt.Errorf("require: Telegram")
	}
	if pf.Vk == nil {
		return fmt.Errorf("require: Vk")
	}
	if pf.Name == nil {
		return fmt.Errorf("require: Name")
	}
	if pf.IsAdmin == nil {
		return fmt.Errorf("require: IsAdmin")
	}
	return nil
}
