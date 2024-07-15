package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateDefaultMedia struct {
	ID int `json:"id"`
	PostDefaultMedia
}

type UpdateDefaultMediaPointer struct {
	ID int `json:"id"`
	PostMediaDefaultPointer
}

func (p *UpdateDefaultMedia) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on UpdateDefaultMedia.Bind: %w", err)
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := UpdateDefaultMediaPointer{}

	err = decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateDefaultMedia.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on UpdateDefaultMedia.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*p = UpdateDefaultMedia{}
	p.Name = *pf.Name
	p.Data = pf.Data
	p.ID = id

	return nil
}
