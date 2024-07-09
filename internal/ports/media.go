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

	// r.Get("/", h.r.Wrap(h.GetAllMedia))
	// r.Get("/{id}", h.r.Wrap(h.GetMedia))
	r.Post("/public", h.r.Wrap(h.PostMediaPublic))
	r.Post("/private", h.r.Wrap(h.PostMediaPrivate))
	// r.Delete("/{id}", h.r.Wrap(h.DeleteMedia))
	// r.Put("/{id}", h.r.Wrap(h.UpdateMedia))

	return r
}

// func (h *MediaHandler) GetAllMedia(w http.ResponseWriter, req *http.Request) handler.Response {
// 	filter := &requests.GetAllMedia{}

// 	res, err := h.media.GetAllMedia(context.Background(), *filter)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.GetAllMedia GetAllMedia")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(res)
// }

// func (h *MediaHandler) GetMedia(w http.ResponseWriter, req *http.Request) handler.Response {
// 	mediaId := &requests.GetMedia{}

// 	err := mediaId.Bind(req)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.GetMedia GetMedia")
// 		return handler.BadRequestResponse()
// 	}

// 	res, err := h.media.GetMedia(context.Background(), mediaId.ID)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.GetMedia GetMedia")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(res)
// }

func (h *MediaHandler) PostMediaPublic(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("PostHandler: got PostMediaPublic request")

	media := &requests.PostMedia{}

	err := media.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaPublic: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("PostHandler: parsed PostMediaPublic request: %v", media)

	_, err = h.media.PostObject(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPublic: %v", err)
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

	_, err = h.media.PostObjectBcrypt(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPrivate: %v", err)
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("PostHandler: done PostMediaPrivate request")

	return handler.OkResponse(nil)
}

// func (h *MediaHandler) DeleteMedia(w http.ResponseWriter, req *http.Request) handler.Response {
// 	media := &requests.DeleteMedia{}

// 	err := media.Bind(req)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.DeleteMedia DeleteMedia")
// 		return handler.BadRequestResponse()
// 	}

// 	err = h.media.DeleteMedia(context.Background(), media.ID)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.DeleteMedia DeleteMedia")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(nil)
// }

// func (h *MediaHandler) UpdateMedia(w http.ResponseWriter, req *http.Request) handler.Response {
// 	media := &requests.UpdateMedia{}

// 	err := media.Bind(req)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.UpdateMedia UpdateMedia")
// 		return handler.BadRequestResponse()
// 	}

// 	err = h.media.UpdateMedia(context.Background(), *mapper.MakeRequestPutMedia(*media))
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.UpdateMedia UpdateMedia")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(nil)
// }
