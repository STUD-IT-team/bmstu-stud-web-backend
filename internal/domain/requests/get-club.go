package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetClub struct {
	ID int `json:"club_id"`
}

func (c *GetClub) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi on club_id in request: %w", err)
	}

	c.ID = id
	return c.validate()
}

func (c *GetClub) validate() error {
	if c.ID == 0 {
		return fmt.Errorf("require: id")
	}
	return nil
}
