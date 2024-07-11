package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type MediaHandler struct {
	r      handler.Renderer
	media  app.MediaService
	logger *log.Logger
	guard  *app.GuardService
}

func NewMediaHandler(r handler.Renderer, media app.MediaService, logger *log.Logger, guard *app.GuardService) *MediaHandler {
	return &MediaHandler{
		r:      r,
		media:  media,
		logger: logger,
		guard:  guard,
	}
}

func (h *MediaHandler) BasePrefix() string {
	return "/media"
}

func (h *MediaHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/public", h.r.Wrap(h.PostMediaPublic))
	r.Post("/private", h.r.Wrap(h.PostMediaPrivate))

	return r
}

func (h *MediaHandler) PostMediaPublic(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("PostHandler: got PostMediaPublic request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token PostMediaPublic: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("Unauthorized request: %v", err)
		return handler.UnauthorizedResponse()
	}

	media := &requests.PostMedia{}

	err = media.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaPublic: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("PostHandler: parsed PostMediaPublic request")

	response, err := h.media.PostMedia(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPublic: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("PostHandler: done PostMediaPublic request")

	return handler.OkResponse(response)
}

func (h *MediaHandler) PostMediaPrivate(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("PostHandler: got PostMediaPrivate request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token PostMediaPublic: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("Unauthorized request: %v", err)
		return handler.UnauthorizedResponse()
	}

	media := &requests.PostMedia{}

	err = media.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaPrivate: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("PostHandler: parsed PostMediaPrivate request")

	response, err := h.media.PostMediaBcrypt(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPrivate: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("PostHandler: done PostMediaPrivate request")

	return handler.OkResponse(response)
}
