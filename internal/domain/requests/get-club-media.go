package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetClubMedia struct {
	ID int `json:"club_id"`
}

func (c *GetClubMedia) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi on club_id: %w", err)
	}
	c.ID = id
	return c.validate()
}

func (c *GetClubMedia) validate() error {
	if c.ID == 0 {
		return fmt.Errorf("require: club_id")
	}
	return nil
}
