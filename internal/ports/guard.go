package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

type GuardHandler struct {
	r      handler.Renderer
	guard  app.GuardService
	logger *log.Logger
}

func NewGuardHandler(r handler.Renderer, guard app.GuardService, logger *log.Logger) *GuardHandler {
	return &GuardHandler{
		r:      r,
		guard:  guard,
		logger: logger,
	}
}

func (h *GuardHandler) BasePrefix() string {
	return "/guard"
}

func (h *GuardHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", h.r.Wrap(h.LoginUser))
	r.Post("/logout", h.r.Wrap(h.LogoutUser))

	return r
}

func (h *GuardHandler) LoginUser(w http.ResponseWriter, req *http.Request) handler.Response {

	h.logger.Infof("GuardHandler: got LoginUser request")
	lreq := &requests.LoginRequest{}

	err := lreq.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request LoginUser: %v", err)
		return handler.BadRequestResponse()
	}

	res, err := h.guard.Login(context.Background(), lreq)
	if err != nil {
		h.logger.Warnf("can't service.LoginUser LoginUser: %v", err)
		return handler.InternalServerErrorResponse()
	}

	h.logger.Infof("GuardHandler: request LoginUser done")

	resp := handler.OkResponse(nil)
	resp.SetKVHeader("Set-Cookie", "AccessToken="+res.AccessToken)

	return resp
}

func (h *GuardHandler) LogoutUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// lreq := &requests.LogoutRequest{}

	// err := lreq.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.LoginUser LoginUser")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.guard.Login(context.Background(), lreq)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.LogoutUser LogoutUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}
