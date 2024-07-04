package requests

import (
	"github.com/go-chi/chi"

	"fmt"
	"net/http"
	"strconv"
)

type DeleteMember struct {
	ID int `json:"id"`
}

func (f *DeleteMember) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on DeleteMember.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *DeleteMember) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}

	return nil
}
