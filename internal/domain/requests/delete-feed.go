package requests

import (
	"github.com/go-chi/chi"

	"fmt"
	"net/http"
	"strconv"
)

type DeleteFeed struct {
	ID int `json:"id"`
}

func (f *DeleteFeed) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on DeleteFeed.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *DeleteFeed) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}

	return nil
}
