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
}

func NewMediaHandler(r handler.Renderer, media app.MediaService, logger *log.Logger) *MediaHandler {
	return &MediaHandler{
		r:      r,
		media:  media,
		logger: logger,
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

	media := &requests.PostMedia{}

	err := media.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaPublic: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("PostHandler: parsed PostMediaPublic request: %v", media)

	_, err = h.media.PostMedia(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPublic: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("PostHandler: done PostMediaPublic request")

	return handler.OkResponse(nil)
}

func (h *MediaHandler) PostMediaPrivate(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("PostHandler: got PostMediaPrivate request")

	media := &requests.PostMedia{}

	err := media.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaPrivate: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("PostHandler: parsed PostMediaPrivate request: %v", media)

	_, err = h.media.PostMediaBcrypt(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPrivate: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("PostHandler: done PostMediaPrivate request")

	return handler.OkResponse(nil)
}
