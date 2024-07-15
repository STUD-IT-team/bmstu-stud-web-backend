package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetDocumentsByCategory struct {
	ID int `json:"id"`
}

func (f *GetDocumentsByCategory) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "category_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on GetDocumentsByCategory.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *GetDocumentsByCategory) validate() error {
	return nil
}
