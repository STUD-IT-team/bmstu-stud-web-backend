package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetFeedEncounters struct {
	ID int `json:"id"`
}

func (f *GetFeedEncounters) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on GetFeedEncounters.Bind: %w", err)
	}

	f.ID = id

	return nil
}
