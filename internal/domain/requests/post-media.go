package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostMedia struct {
	Name string
	Data []byte
}

type PostMediaPointer struct {
	Name *string
	Data []byte
}

func (m *PostMedia) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pm := PostMediaPointer{}

	err := decoder.Decode(&pm)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostMedia.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostMedia.Bind")
	}

	err = pm.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*m = PostMedia{
		Name: *pm.Name,
		Data: pm.Data,
	}

	return m.validate()
}

func (f *PostMedia) validate() error {
	return nil
}

func (pf *PostMediaPointer) validate() error {
	if pf.Data == nil {
		return fmt.Errorf("require: Data")
	}
	if pf.Name == nil {
		return fmt.Errorf("require: Name")
	}
	return nil
}
