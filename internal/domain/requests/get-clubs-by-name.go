package requests

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type GetClubsByName struct {
	Name string `json:"name"`
}

func (c *GetClubsByName) Bind(req *http.Request) error {
	name := chi.URLParam(req, "name")
	c.Name = name
	return c.validate()
}

func (c *GetClubsByName) validate() error {
	if c.Name == "" {
		return fmt.Errorf("require: name - must not be empty")
	}
	return nil
}
