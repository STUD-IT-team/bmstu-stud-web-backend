package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

type RoutesCfg struct {
	BasePath string
}

type Route interface {
	Routes() chi.Router
	BasePrefix() string
}

type GroupHandler struct {
	prefix string
	routes []Route
	router chi.Router
}

func NewGroupHandler(prefix string, rts ...Route) *GroupHandler {
	return &GroupHandler{
		prefix: prefix,
		routes: rts,
		router: chi.NewRouter(),
	}
}

func (h *GroupHandler) Use(middlewares ...func(http.Handler) http.Handler) {
	h.router.Use(middlewares...)
}

func (h *GroupHandler) BasePrefix() string {
	return h.prefix
}

func (h *GroupHandler) Routes() chi.Router {
	router := h.router

	for _, r := range h.routes {
		router.Mount(r.BasePrefix(), r.Routes())
	}

	return router
}

func MakePublicRoutes(
	router *chi.Mux,
	cfg RoutesCfg,
	routes ...Route,
) Option {
	return func(r *Handler) {
		api := r.Group(
			WithRealIP(),
			WithRecover(),
		)

		for i := range routes {
			router.Mount(routes[i].BasePrefix(), routes[i].Routes())
		}

		api.Route(cfg.BasePath, func(r chi.Router) {
			r.Mount("/", router)
		})
	}
}
