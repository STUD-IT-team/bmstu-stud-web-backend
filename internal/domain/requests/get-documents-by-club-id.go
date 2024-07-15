package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetDocumentsByClubID struct {
	ID int `json:"id"`
}

func (f *GetDocumentsByClubID) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on GetDocumentsByClubID.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *GetDocumentsByClubID) validate() error {
	return nil
}
