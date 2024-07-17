package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Telegram string `json:"telegram"`
	Vk       string `json:"vk"`
}

type RegisterPointer struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	Telegram *string `json:"telegram"`
	Vk       *string `json:"vk"`
}

func (r *Register) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pr := RegisterPointer{}

	err := decoder.Decode(&pr)
	if err != nil {
		return fmt.Errorf("can't json decoder on Register.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on Register.Bind")
	}

	err = pr.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	r.Login = *pr.Login
	r.Password = *pr.Password
	r.Name = *pr.Name
	r.Telegram = *pr.Telegram
	r.Vk = *pr.Vk

	return r.validate()
}

func (r *Register) validate() error {
	return nil
}

func (pr *RegisterPointer) validate() error {
	if pr.Login == nil {
		return fmt.Errorf("require: login")
	}
	if pr.Password == nil {
		return fmt.Errorf("require: password")
	}
	if pr.Name == nil {
		return fmt.Errorf("require: name")
	}
	if pr.Telegram == nil {
		return fmt.Errorf("require: telegram")
	}
	if pr.Vk == nil {
		return fmt.Errorf("require: vk")
	}
	return nil
}
