package requests

import (
	"net/http"

	"github.com/go-chi/chi"
)

type GetFeedByTitle struct {
	Search string `json:"title"`
}

func (f *GetFeedByTitle) Bind(req *http.Request) error {
	search := chi.URLParam(req, "title")

	f.Search = search

	return nil
}
