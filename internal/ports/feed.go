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

type FeedHandler struct {
	r    handler.Renderer
	feed app.FeedService
}

func NewFeedHandler(r handler.Renderer, feed app.FeedService) *FeedHandler {
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
	r.Get("/encounters/", h.r.Wrap(h.GetAllFeedEncounters))
	r.Get("/{type}", h.r.Wrap(h.GetFeedByType))
	r.Post("/{feed}", h.r.Wrap(h.PostFeed))
	r.Delete("/{id}", h.r.Wrap(h.DeleteFeed))
	r.Put("/{id}", h.r.Wrap(h.UpdateFeed))

	return r
}

func (h *FeedHandler) GetAllFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	filter := &requests.GetFeedByFilter{}

	err := filter.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetAllFeed GetAllFeed")
		return handler.BadRequestResponse()
	}

	res, err := h.feed.GetFeedByFilter(context.Background(), *filter)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeedByFilter GetFeedByFilter")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *FeedHandler) GetFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	feedId := &requests.GetFeed{}

	err := feedId.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeed GetFeed")
		return handler.BadRequestResponse()
	}

	res, err := h.feed.GetFeed(context.Background(), feedId.ID)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeed GetFeed")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *FeedHandler) GetAllFeedEncounters(w http.ResponseWriter, req *http.Request) handler.Response {
	feedId := &requests.GetFeed{}

	err := feedId.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetAllFeedEncounters GetAllFeedEncounters")
		return handler.BadRequestResponse()
	}

	res, err := h.feed.GetAllFeedEncounters(context.Background())
	if err != nil {
		log.WithError(err).Warnf("can't service.GetAllFeedEncounters GetFeedEncounters")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *FeedHandler) GetFeedByType(w http.ResponseWriter, req *http.Request) handler.Response {
	res, err := h.feed.GetFeedByType(context.Background())
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeedByType GetFeedByType")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *FeedHandler) PostFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	feed := &requests.PostFeed{}

	err := feed.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.PostFeed PostFeed")
		return handler.BadRequestResponse()
	}

	err = h.feed.PostFeed(context.Background(), *mapper.MakeRequestPutFeed(*feed))
	if err != nil {
		log.WithError(err).Warnf("can't service.PostFeed PostFeed")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(nil)
}

func (h *FeedHandler) DeleteFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	feedId := &requests.DeleteFeed{}

	err := feedId.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.DeleteFeed DeleteFeed")
		return handler.BadRequestResponse()
	}

	err = h.feed.DeleteFeed(context.Background(), feedId.ID)
	if err != nil {
		log.WithError(err).Warnf("can't service.DeleteFeed DeleteFeed")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(nil)
}

func (h *FeedHandler) UpdateFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	feed := &requests.UpdateFeed{}

	err := feed.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.UpdateFeed UpdateFeed")
		return handler.BadRequestResponse()
	}

	err = h.feed.UpdateFeed(context.Background(), *mapper.MakeRequestPutFeed(*feed))
	if err != nil {
		log.WithError(err).Warnf("can't service.UpdateFeed UpdateFeed")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(nil)
}
