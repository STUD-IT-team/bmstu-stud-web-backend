package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
)

type APIHandler struct {
	r   handler.Renderer
	api app.API
}

func NewAPIHandler(r handler.Renderer, api app.API) *APIHandler {
	return &APIHandler{
		r:   r,
		api: api,
	}
}

func (h *APIHandler) BasePrefix() string {
	return "/api"
}

func (h *APIHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/echo", h.r.Wrap(h.echo))

	return r
}

// echo
//
// @Summary    Echo endpoint
// @Description Возвращает базовую информацию о состоянии API
// @Tags      public.api
// @Produce    json
// @Success    200      {object}  responses.GetEcho
// @Router      /api/echo [get]
// @Security    Public
func (h *APIHandler) echo(_ http.ResponseWriter, r *http.Request) handler.Response {
	out := h.api.Echo()

	return handler.OkResponse(out)
}
