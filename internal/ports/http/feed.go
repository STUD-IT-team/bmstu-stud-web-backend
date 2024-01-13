package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/storage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type FeedHandler struct {
	r    handler.Renderer
	feed storage.Storage
}

func NewFeedHandler(r handler.Renderer, feed storage.Storage) *FeedHandler {
	return &FeedHandler{
		r:    r,
		feed: feed,
	}
}

func (h *FeedHandler) BasePrefix() string {
	return "/feed"
}

func (h *FeedHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllFeed))
	r.Get("/{id}", h.r.Wrap(h.GetFeed))

	return r
}

func (h *FeedHandler) GetAllFeed(w http.ResponseWriter, _ *http.Request) handler.Response {
	res, err := h.feed.GetAllFeed(context.Background())
	if err != nil {
		log.WithError(err).Warnf("can't service.GetAllFeed GetAllFeed")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *FeedHandler) GetFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	id, err := requests.NewGetFeed().Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeed GetFeed")
		return handler.BadRequestResponse()
	}

	res, err := h.feed.GetFeed(context.Background(), id)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeed GetFeed")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}
