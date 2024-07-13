package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type FeedHandler struct {
	r      handler.Renderer
	feed   app.FeedService
	logger *log.Logger
	guard  *app.GuardService
}

func NewFeedHandler(r handler.Renderer, feed app.FeedService, logger *log.Logger, guard *app.GuardService) *FeedHandler {
	return &FeedHandler{
		r:      r,
		feed:   feed,
		logger: logger,
		guard:  guard,
	}
}

func (h *FeedHandler) BasePrefix() string {
	return "/feed"
}

func (h *FeedHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllFeed))
	r.Get("/{id}", h.r.Wrap(h.GetFeed))
	r.Get("/encounters/{id}", h.r.Wrap(h.GetFeedEncounters))
	r.Get("/search/{title}", h.r.Wrap(h.GetFeedByTitle))
	r.Post("/", h.r.Wrap(h.PostFeed))
	r.Delete("/{id}", h.r.Wrap(h.DeleteFeed))
	r.Put("/{id}", h.r.Wrap(h.UpdateFeed))

	return r
}

// GetAllFeed retrieves all feed items
//
//	@Summary     Retrieve all feed items
//	@Description Get a list of all feed items
//	@Tags        public.feed
//	@Produce     json
//	@Success     200 {object}  responses.GetAllFeed
//	@Failure     404
//	@Router      /feed [get]
//	@Security    public
func (h *FeedHandler) GetAllFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FeedHandler: got GetAllFeed request")

	res, err := h.feed.GetAllFeed(context.Background())
	if err != nil {
		h.logger.Warnf("can't FeedService.GetAllFeed: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("FeedHandler: request GetAllFeed done")

	return handler.OkResponse(res)
}

// GetFeed retrieves a feed item by its ID
//
//	@Summary     Retrieve feed item by ID
//	@Description Get a specific feed item using its ID
//	@Tags        public.feed
//	@Produce     json
//	@Param       id   path     string           true "id"
//	@Success     200  {object} responses.GetFeed
//	@Failure     400
//	@Failure     404
//	@Router      /feed/{id} [get]
//	@Security    public
func (h *FeedHandler) GetFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FeedHandler: got GetFeed request")

	feedId := &requests.GetFeed{}

	err := feedId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("FeedHandler: parse request GetFeed: %v", feedId)

	res, err := h.feed.GetFeed(context.Background(), feedId.ID)
	if err != nil {
		h.logger.Warnf("can't FeedService.GetFeed: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("FeedHandler: request GetFeed done")

	return handler.OkResponse(res)
}

// GetFeedEncounters retrieves by id
//
//	@Summary     Retrieve encounters by ID
//	@Description Get encounters using ID (0 for main page)
//	@Tags        public.feed
//	@Produce     json
//	@Param       id   path     string           true "id"
//	@Success     200  {object} responses.GetFeedEncounters
//	@Failure     400
//	@Failure     404
//	@Router      /feed/encounters/{id} [get]
//	@Security    public
func (h *FeedHandler) GetFeedEncounters(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FeedHandler: got GetFeedEncounters request")

	encounterId := &requests.GetFeedEncounters{}

	err := encounterId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetFeedEncounters: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("FeedHandler: parse request GetFeedEncounters: %v", encounterId)

	res, err := h.feed.GetFeedEncounters(context.Background(), encounterId.ID)
	if err != nil {
		h.logger.Warnf("can't FeedService.GetFeedEncounters: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("FeedHandler: request GetFeedEncounters done")

	return handler.OkResponse(res)
}

// GetFeedByTitle retrieves feed items by title
//
//	@Summary     Retrieve feed items by title
//	@Description Get feed items that match the specified title
//	@Tags        public.feed
//	@Produce     json
//	@Param       title path    string           true "title"
//	@Success     200  {object} responses.GetFeedByTitle
//	@Failure     400
//	@Failure     404
//	@Router      /feed/search/{title} [get]
//	@Security    public
func (h *FeedHandler) GetFeedByTitle(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FeedHandler: got GetFeedByTitle request")

	filter := &requests.GetFeedByTitle{}

	err := filter.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetFeedByTitle: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("FeedHandler: parse request GetFeedByTitle: %v", filter)

	res, err := h.feed.GetFeedByTitle(context.Background(), *filter)
	if err != nil {
		h.logger.Warnf("can't FeedService.GetFeedByTitle: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("FeedHandler: request GetFeedByTitle done")

	return handler.OkResponse(res)
}

// PostFeed creates a new feed item
//
//		@Summary     Create a new feed item
//		@Description Create a new feed item with the provided data
//		@Tags        auth.feed
//		@Accept      json
//		@Param       request body requests.PostFeed true "Feed data"
//		@Success     201
//		@Failure     400
//	 	@Failure     401
//		@Failure     500
//		@Router      /feed [post]
//		@Security    Authorised
func (h *FeedHandler) PostFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FeedHandler: got PostFeed request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token PostFeed: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on PostFeed: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("FeedHandler: PostFeed Authenticated: %v", resp.MemberID)

	feed := &requests.PostFeed{}

	err = feed.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind PostFeed: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("FeedHandler: parse request PostFeed: %v", feed)

	err = h.feed.PostFeed(context.Background(), mapper.MakeRequestPostFeed(feed))
	if err != nil {
		h.logger.Warnf("can't FeedService.PostFeed: %v", err)
		if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("FeedHandler: request PostFeed done")

	return handler.CreatedResponse(nil)
}

// DeleteFeed deletes a feed item by ID
//
//	@Summary     Delete a feed item by ID
//	@Description Delete a specific feed item using its ID
//	@Tags        auth.feed
//	@Param       id   path     string           true "Feed ID"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Router      /feed/{id} [delete]
//	@Security    Authorised
func (h *FeedHandler) DeleteFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FeedHandler: got DeleteFeed request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token DeleteFeed: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteFeed: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("FeedHandler: DeleteFeed Authenticated: %v", resp.MemberID)

	feedId := &requests.DeleteFeed{}

	err = feedId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind DeleteFeed: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("FeedHandler: parse request DeleteFeed: %v", feedId)

	err = h.feed.DeleteFeed(context.Background(), feedId.ID)
	if err != nil {
		h.logger.Warnf("can't FeedService.DeleteFeed: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("FeedHandler: request DeleteFeed done")

	return handler.OkResponse(nil)
}

// UpdateFeed updates a feed item
//
//	@Summary     Update a feed item
//	@Description Update an existing feed item with the provided data
//	@Tags        auth.feed
//	@Accept      json
//	@Param       id   path     string           true "Feed ID"
//	@Param       request body requests.PostFeed true "Feed new data"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     500
//	@Router      /feed/{id} [put]
//	@Security    Authorised
func (h *FeedHandler) UpdateFeed(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FeedHandler: got UpdateFeed request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token UpdateFeed: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on UpdateFeed: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("FeedHandler: UpdateFeed Authenticated: %v", resp.MemberID)

	feed := &requests.UpdateFeed{}

	err = feed.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind UpdateFeed: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("FeedHandler: parse request UpdateFeed: %v", feed)

	err = h.feed.UpdateFeed(context.Background(), mapper.MakeRequestUpdateFeed(feed))
	if err != nil {
		h.logger.Warnf("can't FeedService.UpdateFeed: %v", err)
		if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("FeedHandler: request UpdateFeed done")

	return handler.OkResponse(nil)
}
