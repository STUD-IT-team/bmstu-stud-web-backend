package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
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
	r.Delete("/{id}", h.r.Wrap(h.DeleteFeed))

	return r
}

// GetAllFeed get all feeds
//
//	@Summary      List feeds
//	@Description  get feeds
//	@Tags         feed
//	@Accept       json
//	@Produce      json
//	@Success      200  {array}   responses.GetAllFeed
//	@Failure      500  {object}  handler.Response
//	@Router       /feed [get]
func (h *FeedHandler) GetAllFeed(w http.ResponseWriter, _ *http.Request) handler.Response {
	res, err := h.feed.GetAllFeed(context.Background())
	if err != nil {
		log.WithError(err).Warnf("can't service.GetAllFeed GetAllFeed")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

// GetFeed get feed by id
//
//	@Summary      feed
//	@Description  get feed by id
//	@Tags         feed
//	@Accept       json
//	@Produce      json
//	@Param        id path string true "feed ID"
//	@Success      200  {object}  responses.GetFeed
//	@Failure      400  {object}  handler.Response
//	@Failure      500  {object}  handler.Response
//	@Router       /feed/{id} [get]
func (h *FeedHandler) GetFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	feedId := requests.NewGetFeed()
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

// DeleteFeed delete feed by id
//
//	@Summary      feed
//	@Description  delete feed by id
//	@Tags         feed
//	@Accept       json
//	@Produce      json
//	@Param        id path string true "feed ID"
//	@Success      200  {object}  handler.Response
//	@Failure      400  {object}  handler.Response
//	@Failure      500  {object}  handler.Response
//	@Router       /feed/{id} [delete]
func (h *FeedHandler) DeleteFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	feedId := requests.NewDeleteFeed()
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
