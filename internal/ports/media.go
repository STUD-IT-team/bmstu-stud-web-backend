package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type MediaHandler struct {
	r     handler.Renderer
	media app.MediaService
}

func NewMediaHandler(r handler.Renderer, media app.MediaService) *MediaHandler {
	return &MediaHandler{
		r:     r,
		media: media,
	}
}

func (h *MediaHandler) BasePrefix() string {
	return "/media"
}

func (h *MediaHandler) Routes() chi.Router {
	r := chi.NewRouter()

	// r.Get("/", h.r.Wrap(h.GetAllMedia))
	// r.Get("/{id}", h.r.Wrap(h.GetMedia))
	r.Post("/", h.r.Wrap(h.PostMedia)) // TODO: redoo
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

func (h *MediaHandler) PostMedia(w http.ResponseWriter, req *http.Request) handler.Response {
	media := &requests.PostMedia{}

	err := media.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.PostMedia PostMedia")
		return handler.BadRequestResponse()
	}

	err = h.media.PostMedia(context.Background(), *mapper.MakeRequestPutMedia(*media))
	if err != nil {
		log.WithError(err).Warnf("can't service.PostMedia PostMedia")
		return handler.InternalServerErrorResponse()
	}

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
