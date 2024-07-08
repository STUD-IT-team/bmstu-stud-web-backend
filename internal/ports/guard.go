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

	access, err := getAccessToken(req)
	if err == nil {
		h.logger.Infof("GuardHandler: Access token found: %v", access)

		resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: access})
		if err == nil && resp.Valid {
			h.logger.Infof("GuardHandler: User already authenticated: %v", resp.MemberID)
			return handler.OkResponse(nil)
		}
	}

	err = lreq.Bind(req)
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
	resp.SetKVHeader("Set-Cookie", "AccessToken="+res.AccessToken+"; Path=/; HttpOnly")

	return resp
}

func (h *GuardHandler) LogoutUser(w http.ResponseWriter, req *http.Request) handler.Response {

	h.logger.Infof("GuardHandler: got LogoutUser request")

	access, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token LogoutUser: %v", err)
		return handler.UnauthorizedResponse()
	}

	err = h.guard.Logout(context.Background(), &requests.LogoutRequest{AccessToken: access})
	if err != nil {
		h.logger.Warnf("can't service.LogoutUser LogoutUser: %v", err)
		return handler.InternalServerErrorResponse()
	}

	// Можн дописать удаление куки(проставить expireat=unix(0))

	h.logger.Infof("GuardHandler: request LogoutUser done")
	return handler.OkResponse(nil)
}
