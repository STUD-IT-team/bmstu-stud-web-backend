package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetFeed struct {
	ID int `json:"id"`
}

func NewGetFeed() *GetFeed {
	return &GetFeed{}
}

func (f *GetFeed) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	f.ID = id
	return err
}
