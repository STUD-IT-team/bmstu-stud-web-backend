package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type GuardHandler struct {
	r     handler.Renderer
	guard app.GuardService
}

func NewGuardHandler(r handler.Renderer, guard app.GuardService) *GuardHandler {
	return &GuardHandler{
		r:     r,
		guard: guard,
	}
}

func (h *GuardHandler) BasePrefix() string {
	return "/guard"
}

func (h *GuardHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login/{user}", h.r.Wrap(h.LoginUser))
	r.Post("/logout/{user}", h.r.Wrap(h.LogoutUser))

	return r
}

func (h *GuardHandler) LoginUser(w http.ResponseWriter, req *http.Request) handler.Response {
	lreq := &requests.LoginRequest{}

	err := lreq.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.LoginUser LoginUser")
		return handler.BadRequestResponse()
	}

	res, err := h.guard.Login(context.Background(), lreq)
	if err != nil {
		log.WithError(err).Warnf("can't service.LoginUser LoginUser")
		return handler.InternalServerErrorResponse()
	}

	return res
}

func (h *GuardHandler) LogoutUser(w http.ResponseWriter, req *http.Request) handler.Response {
	lreq := &requests.LogoutRequest{}

	err := lreq.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.LoginUser LoginUser")
		return handler.BadRequestResponse()
	}

	res, err := h.guard.Login(context.Background(), lreq)
	if err != nil {
		log.WithError(err).Warnf("can't service.LogoutUser LogoutUser")
		return handler.InternalServerErrorResponse()
	}

	return res
}
