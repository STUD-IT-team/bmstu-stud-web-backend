package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequestPointer struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
}

func (lr *LoginRequest) Bind(req *http.Request) error {
	pl := LoginRequestPointer{}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&pl)
	if err != nil {
		return err
	}

	err = pl.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*lr = LoginRequest{
		Login:    *pl.Login,
		Password: *pl.Password,
	}
	return lr.validate()
}

func (pl *LoginRequestPointer) validate() error {
	if pl.Login == nil {
		return fmt.Errorf("require: login")
	}
	if pl.Password == nil {
		return fmt.Errorf("require: password")
	}
	return nil
}

func (pl *LoginRequest) validate() error {
	return nil
}
