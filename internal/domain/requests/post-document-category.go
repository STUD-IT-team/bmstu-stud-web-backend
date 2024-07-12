package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostDocumentCategory struct {
	Name string `json:"name"`
}

type PostDocumentCategoryPointer struct {
	Name *string `json:"name"`
}

func (v *PostDocumentCategory) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pv := PostDocumentCategoryPointer{}

	err := decoder.Decode(&pv)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostDocumentCategory.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostDocumentCategory.Bind")
	}

	err = pv.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*v = PostDocumentCategory{Name: *pv.Name}

	return v.validate()
}

func (v *PostDocumentCategory) validate() error {
	return nil
}

func (pv *PostDocumentCategoryPointer) validate() error {
	if pv.Name == nil {
		return fmt.Errorf("require: Name")
	}
	return nil
}
