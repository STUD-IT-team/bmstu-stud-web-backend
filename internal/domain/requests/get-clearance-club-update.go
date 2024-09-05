package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetClearanceClubUpdate struct {
	ClubID int `json:"club_id"`
}

func (p *GetClearanceClubUpdate) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on DeleteMember.Bind: %w", err)
	}

	p.ClubID = id

	return p.validate()
}

func (p *GetClearanceClubUpdate) validate() error {
	return nil
}
