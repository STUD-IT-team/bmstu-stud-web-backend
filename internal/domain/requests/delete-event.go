package requests

import (
	"github.com/go-chi/chi"

	"fmt"
	"net/http"
	"strconv"
)

type DeleteEvent struct {
	ID int `json:"id"`
}

func (f *DeleteEvent) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on DeleteEvent.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *DeleteEvent) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}

	return nil
}
