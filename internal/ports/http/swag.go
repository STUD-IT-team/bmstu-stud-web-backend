package http

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
)

type SwagHandler struct {
	r handler.Renderer
}

func NewSwagHandler(r handler.Renderer) *SwagHandler {
	return &SwagHandler{
		r: r,
	}
}

func (h *SwagHandler) BasePrefix() string {
	return "/docs"
}

func (h *SwagHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", httpSwagger.Handler())

	return r
}
