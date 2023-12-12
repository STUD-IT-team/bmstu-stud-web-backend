package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func (h *APIHandler) GetAllFeed(w http.ResponseWriter, _ *http.Request) handler.Response {
	res, err := h.feed.GetAllFeed(context.Background())

	if err != nil {
		log.WithField("", "GetAllFeed").Error(err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *APIHandler) GetFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		log.WithField("", "GetFeed").Error(err)
		return handler.BadRequestResponse()
	}

	res, err := h.feed.GetFeed(context.Background(), id)

	if err != nil {
		log.WithField("", "GetFeed").Error(err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}
