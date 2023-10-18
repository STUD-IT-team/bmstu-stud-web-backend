package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Handler represents HTTP route multiplexer
type Handler struct {
	*chi.Mux
}

// New creates a new http.Handler with the provided options
func New(opts ...Option) *Handler {
	r := &Handler{Mux: chi.NewRouter()}
	r.Use(render.SetContentType(render.ContentTypeJSON))

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *Handler) Group(opts ...Option) *Handler {
	in := r.Mux.With().(*chi.Mux)
	h := &Handler{Mux: in}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

// Option represents Handler option
type Option func(*Handler)

// WithRealIP adds middleware which helps get real requester's IP, not proxy
func WithRealIP() Option {
	return func(r *Handler) { r.Use(middleware.RealIP) }
}

// WithRecover adds recover middleware, which can catch panics from handlers
func WithRecover() Option {
	return func(r *Handler) { r.Use(middleware.Recoverer) }
}
