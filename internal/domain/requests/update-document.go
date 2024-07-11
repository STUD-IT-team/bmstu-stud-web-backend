package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateDocument struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Data   []byte `json:"data"`
	ClubID int    `json:"club_id"`
}

type UpdateDocumentPointer struct {
	Name   *string `json:"name"`
	Data   []byte  `json:"data"`
	ClubID *int    `json:"club_id"`
}

func (v *UpdateDocument) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pv := UpdateDocumentPointer{}

	err := decoder.Decode(&pv)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateDocument.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on UpdateDocument.Bind")
	}

	err = pv.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*v = UpdateDocument{
		Name:   *pv.Name,
		Data:   pv.Data,
		ClubID: *pv.ClubID,
	}

	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on UpdateDocument.Bind: %w", err)
	}

	v.ID = id

	return v.validate()
}

func (v *UpdateDocument) validate() error {
	if v.ID == 0 {
		return fmt.Errorf("require: id")
	}
	return nil
}

func (pv *UpdateDocumentPointer) validate() error {
	if pv.Name == nil {
		return fmt.Errorf("require: Name")
	}
	if pv.Data == nil {
		return fmt.Errorf("require: Data")
	}
	if pv.ClubID == nil {
		return fmt.Errorf("require: ClubID")
	}
	return nil
}
