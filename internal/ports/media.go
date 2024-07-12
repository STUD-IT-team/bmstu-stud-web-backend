package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type MediaHandler struct {
	r      handler.Renderer
	media  *app.MediaService
	logger *log.Logger
	guard  *app.GuardService
}

func NewMediaHandler(r handler.Renderer, media *app.MediaService, logger *log.Logger, guard *app.GuardService) *MediaHandler {
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
	r.Get("/default", h.r.Wrap(h.GetMediaDefault))
	r.Get("/default/{id}", h.r.Wrap(h.GetMediaDefaultByID))
	r.Post("/default", h.r.Wrap(h.PostMediaDefault))
	r.Delete("/default", h.r.Wrap(h.DeleteMediaDefault))
	r.Put("/default/{id}", h.r.Wrap(h.UpdateMediaDefault))

	return r
}

// PostMediaPublic
//
// @Summary    Загружает изображение в базу данных
// @Description Загружает изображение в базу данных публично, то есть в хранилище хранится файл по тому же названию, что и подан на вход.
// @Tags      auth.media
// @Produce    json
// @Param      request  body    requests.PostMedia  true  "post media data"
// @Success    200   {object}  responses.PostMedia
// @Failure    400
// @Failure    401
// @Failure    500
// @Router      /media/public [post]
// @Security    Authorized
func (h *MediaHandler) PostMediaPublic(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MediaHandler: got PostMediaPublic request")

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

	h.logger.Infof("MediaHandler: authenticated: %v", resp.MemberID)

	media := &requests.PostMedia{}

	err = media.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaPublic: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MediaHandler: parsed PostMediaPublic request")

	response, err := h.media.PostMedia(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPublic: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("MediaHandler: done PostMediaPublic request")

	return handler.OkResponse(response)
}

// PostMediaPrivate
//
// @Summary    Загружает изображение в базу данных
// @Description Загружает изображение в базу данных приватно, то есть название загруженного файла и хранящегося объетка различаются. По сути из вне нельзя заранее узнать ключ для получения файла.
// @Tags      auth.media
// @Produce    json
// @Param      request  body    requests.PostMedia  true  "post media data"
// @Success    200   {object}  responses.PostMedia
// @Failure    400
// @Failure    401
// @Failure    500
// @Router      /media/private [post]
// @Security    Authorized
func (h *MediaHandler) PostMediaPrivate(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MediaHandler: got PostMediaPrivate request")

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

	h.logger.Infof("MediaHandler: authenticated: %v", resp.MemberID)

	media := &requests.PostMedia{}

	err = media.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaPrivate: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MediaHandler: parsed PostMediaPrivate request")

	response, err := h.media.PostMediaBcrypt(context.Background(), media)
	if err != nil {
		h.logger.Warnf("can't service.PostMedia PostMediaPrivate: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("MediaHandler: done PostMediaPrivate request")

	return handler.OkResponse(response)
}

func (h *MediaHandler) GetMediaDefault(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Infof("Mediahandler: got GetMediaDefault request")

	resp, err := h.media.GetAllMediaDefault(context.Background())
	if err != nil {
		h.logger.Warnf("can't service.GetAllMediaDefault GetMediaDefault: %v", err)
		return handler.NotFoundResponse()
	}
	h.logger.Infof("MediaHandler: done GetMediaDefault request")
	return handler.OkResponse(resp)
}

func (h *MediaHandler) GetMediaDefaultByID(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Infof("Mediahandler: got GetMediaDefaultById request")
	defaultMedia := requests.GetDefaultMedia{}

	err := defaultMedia.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request GetMediaDefault: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MediaHandler: parsed GetMediaDefault request")

	response, err := h.media.GetMediaDefault(context.Background(), defaultMedia.ID)

	if err != nil {
		h.logger.Warnf("can't service.GetMediaDefault GetMediaDefault: %v", err)
		return handler.NotFoundResponse()
	}
	h.logger.Infof("MediaHandler: done GetMediaDefault request")
	return handler.OkResponse(response)
}

func (h *MediaHandler) PostMediaDefault(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Infof("Mediahandler: got PostMediaDefault request")

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

	h.logger.Infof("MediaHandler: authenticated: %v", resp.MemberID)

	defaultMedia := requests.PostDefaultMedia{}

	err = defaultMedia.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request PostMediaDefault: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MediaHandler: parsed PostMediaDefault request")

	response, err := h.media.PutMediaDefault(context.Background(), defaultMedia.Name, defaultMedia.Data)
	if err != nil {
		h.logger.Warnf("can't service.PutMediaDefault PostMediaDefault: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		}
		return handler.InternalServerErrorResponse()
	}
	h.logger.Infof("MediaHandler: done PostMediaDefault request")
	return handler.OkResponse(response)
}

func (h *MediaHandler) DeleteMediaDefault(w http.ResponseWriter, req *http.Request) handler.Response {
	return nil
}

func (h *MediaHandler) UpdateMediaDefault(w http.ResponseWriter, req *http.Request) handler.Response {
	return nil
}
