package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetDocument struct {
	ID int `json:"id"`
}

func (f *GetDocument) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on GetDocument.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *GetDocument) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}

	return nil
}
