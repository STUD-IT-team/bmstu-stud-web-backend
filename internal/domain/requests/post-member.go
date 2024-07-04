package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostMember struct {
	HashPassword []byte `db:"hash_password"`
	Login        string `db:"login"`
	MediaID      int    `db:"media_id"`
	Telegram     string `db:"telegram"`
	Vk           string `db:"vk"`
	Name         string `db:"name"`
	RoleID       int    `db:"role_id"`
	IsAdmin      bool   `db:"isAdmin"`
}

type PostMemberPointer struct {
	HashPassword *[]byte `db:"hash_password"`
	Login        *string `db:"login"`
	MediaID      *int    `db:"media_id"`
	Telegram     *string `db:"telegram"`
	Vk           *string `db:"vk"`
	Name         *string `db:"name"`
	RoleID       *int    `db:"role_id"`
	IsAdmin      *bool   `db:"isAdmin"`
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
		HashPassword: *pf.HashPassword,
		Login:        *pf.Login,
		MediaID:      *pf.MediaID,
		Telegram:     *pf.Telegram,
		Vk:           *pf.Vk,
		Name:         *pf.Name,
		RoleID:       *pf.RoleID,
		IsAdmin:      *pf.IsAdmin,
	}

	return f.validate()
}

func (f *PostMember) validate() error {
	return nil
}

func (pf *PostMemberPointer) validate() error {
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
