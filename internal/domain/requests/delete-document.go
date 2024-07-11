package requests

import (
	"github.com/go-chi/chi"

	"fmt"
	"net/http"
	"strconv"
)

type DeleteDocument struct {
	ID int `json:"id"`
}

func (f *DeleteDocument) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on DeleteDocument.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *DeleteDocument) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}

	return nil
}
