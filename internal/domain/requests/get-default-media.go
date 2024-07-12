package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetDefaultMedia struct {
	ID int `json:"id"`
}

func (c *GetDefaultMedia) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi on id in request: %w", err)
	}

	c.ID = id
	return c.validate()
}

func (c *GetDefaultMedia) validate() error {
	if c.ID == 0 {
		return fmt.Errorf("require: id")
	}
	return nil
}
