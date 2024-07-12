package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetDocumentCategory struct {
	ID int `json:"id"`
}

func (f *GetDocumentCategory) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on GetDocumentCategory.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *GetDocumentCategory) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}

	return nil
}
