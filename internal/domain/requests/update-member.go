package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateMember struct {
	ID           int    `json:"id"`
	HashPassword []byte `json:"hash_password"`
	Login        string `json:"login"`
	MediaID      int    `json:"media_id"`
	Telegram     string `json:"telegram"`
	Vk           string `json:"vk"`
	Name         string `json:"name"`
	RoleID       int    `json:"role_id"`
	IsAdmin      bool   `json:"isAdmin"`
}

type UpdateMemberPointer struct {
	HashPassword *[]byte `json:"hash_password"`
	Login        *string `json:"login"`
	MediaID      *int    `json:"media_id"`
	Telegram     *string `json:"telegram"`
	Vk           *string `json:"vk"`
	Name         *string `json:"name"`
	RoleID       *int    `json:"role_id"`
	IsAdmin      *bool   `json:"isAdmin"`
}

func (f *UpdateMember) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := UpdateMemberPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateMember.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on UpdateMember.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*f = UpdateMember{
		HashPassword: *pf.HashPassword,
		Login:        *pf.Login,
		MediaID:      *pf.MediaID,
		Telegram:     *pf.Telegram,
		Vk:           *pf.Vk,
		Name:         *pf.Name,
		RoleID:       *pf.RoleID,
		IsAdmin:      *pf.IsAdmin,
	}

	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on UpdateMember.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *UpdateMember) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}
	return nil
}

func (pf *UpdateMemberPointer) validate() error {
	if pf.HashPassword == nil {
		return fmt.Errorf("require: HashPassword")
	}
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
	if pf.RoleID == nil {
		return fmt.Errorf("require: RoleID")
	}
	if pf.IsAdmin == nil {
		return fmt.Errorf("require: IsAdmin")
	}
	return nil
}
