package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostDocument struct {
	Name   string `db:"name"`
	Key    string `db:"key"`
	ClubID int    `db:"club_id"`
}

type PostDocumentPointer struct {
	Name   *string `db:"name"`
	Key    *string `db:"key"`
	ClubID *int    `db:"club_id"`
}

func (v *PostDocument) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pv := PostDocumentPointer{}

	err := decoder.Decode(&pv)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostDocument.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostDocument.Bind")
	}

	err = pv.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*v = PostDocument{
		Name:   *pv.Name,
		Key:    *pv.Key,
		ClubID: *pv.ClubID,
	}

	return v.validate()
}

func (v *PostDocument) validate() error {
	return nil
}

func (pv *PostDocumentPointer) validate() error {
	if pv.Name == nil {
		return fmt.Errorf("require: Name")
	}
	if pv.Key == nil {
		return fmt.Errorf("require: Key")
	}
	if pv.ClubID == nil {
		return fmt.Errorf("require: ClubID")
	}
	return nil
}
