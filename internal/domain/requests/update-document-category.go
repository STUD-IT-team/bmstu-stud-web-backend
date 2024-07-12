package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateDocumentCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateDocumentCategoryPointer struct {
	Name *string `json:"name"`
}

func (v *UpdateDocumentCategory) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pv := UpdateDocumentCategoryPointer{}

	err := decoder.Decode(&pv)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateDocumentCategory.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on UpdateDocumentCategory.Bind")
	}

	err = pv.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*v = UpdateDocumentCategory{Name: *pv.Name}

	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on UpdateDocumentCategory.Bind: %w", err)
	}

	v.ID = id

	return v.validate()
}

func (v *UpdateDocumentCategory) validate() error {
	return nil
}

func (pv *UpdateDocumentCategoryPointer) validate() error {
	if pv.Name == nil {
		return fmt.Errorf("require: Name")
	}
	return nil
}
