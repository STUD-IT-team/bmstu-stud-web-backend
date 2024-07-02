package requests

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type GetClubsByType struct {
	Type string `json:"type"`
}

func (c *GetClubsByType) Bind(req *http.Request) error {
	typeStr := chi.URLParam(req, "type")
	c.Type = typeStr
	return c.validate()
}

func (c *GetClubsByType) validate() error {
	if c.Type == "" {
		return fmt.Errorf("require: type - not empty")
	}
	return nil
}
