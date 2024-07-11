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

func (f *PostDocument) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := PostDocumentPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostDocument.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostDocument.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*f = PostDocument{
		Name:   *pf.Name,
		Key:    *pf.Key,
		ClubID: *pf.ClubID,
	}

	return f.validate()
}

func (f *PostDocument) validate() error {
	return nil
}

func (pf *PostDocumentPointer) validate() error {
	if pf.Name == nil {
		return fmt.Errorf("require: Name")
	}
	if pf.Key == nil {
		return fmt.Errorf("require: Key")
	}
	if pf.ClubID == nil {
		return fmt.Errorf("require: ClubID")
	}
	return nil
}
