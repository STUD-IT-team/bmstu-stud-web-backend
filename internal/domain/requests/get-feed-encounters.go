package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetFeedEncounters struct {
	ClubID int `json:"club_id"`
}

func (f *GetFeedEncounters) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on GetFeedEncounters.Bind: %w", err)
	}

	f.ClubID = id

	return nil
}
