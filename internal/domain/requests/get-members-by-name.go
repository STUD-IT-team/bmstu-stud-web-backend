package requests

import (
	"net/http"

	"github.com/go-chi/chi"
)

type GetMembersByName struct {
	Search string `json:"name"`
}

func (f *GetMembersByName) Bind(req *http.Request) error {
	search := chi.URLParam(req, "name")

	f.Search = search

	return nil
}
