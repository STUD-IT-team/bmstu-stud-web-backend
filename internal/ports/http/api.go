package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
)

type APIHandler struct {
	r    handler.Renderer
	api  app.API
	feed app.FeedServiceSrorage
}

func NewAPIHandler(r handler.Renderer, api app.API, feed app.FeedServiceSrorage) *APIHandler {
	return &APIHandler{
		r:    r,
		api:  api,
		feed: feed,
	}
}

func (h *APIHandler) BasePrefix() string {
	return "/api"
}

func (h *APIHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/echo", h.r.Wrap(h.echo))

	r.Get("/feed", h.r.Wrap(h.GetAllFeed))
	r.Get("/feed/:id", h.r.Wrap(h.GetFeed))

	return r
}

func (h *APIHandler) echo(_ http.ResponseWriter, r *http.Request) handler.Response {
	out := h.api.Echo()

	return handler.OkResponse(out)
}
