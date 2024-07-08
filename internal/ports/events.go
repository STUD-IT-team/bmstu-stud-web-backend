package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

type EventsHandler struct {
	r      handler.Renderer
	events app.EventsService
	logger *log.Logger
	guard  *app.GuardService
}

func NewEventsHandler(r handler.Renderer, events app.EventsService, logger *log.Logger, guard *app.GuardService) *EventsHandler {
	return &EventsHandler{
		r:      r,
		events: events,
		logger: logger,
		guard:  guard,
	}
}

func (h *EventsHandler) BasePrefix() string {
	return "/events"
}

func (h *EventsHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllEvents))
	r.Get("/{id}", h.r.Wrap(h.GetEvent))
	r.Get("/range/", h.r.Wrap(h.GetEventsByRange))
	r.Post("/", h.r.Wrap(h.PostEvent))
	r.Delete("/{id}", h.r.Wrap(h.DeleteEvent))
	r.Put("/{id}", h.r.Wrap(h.UpdateEvent))

	return r
}

// GetAllEvents retrieves all events items
//
//	@Summary     Retrieve all events items
//	@Description Get a list of all events items
//	@Tags        public.events
//	@Produce     json
//	@Success     200 {object}  responses.GetAllEvents
//	@Failure     404
//	@Router      /events [get]
//	@Security    public
func (h *EventsHandler) GetAllEvents(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("EventsHandler: got GetAllEvents request")

	res, err := h.events.GetAllEvents(context.Background())
	if err != nil {
		h.logger.Warnf("can't EventsService.GetAllEvents: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("EventsHandler: request GetAllEvents done")

	return handler.OkResponse(res)
}

// GetEvent retrieves an event item by its ID
//
//	@Summary     Retrieve event item by ID
//	@Description Get a specific event item using its ID
//	@Tags        public.events
//	@Produce     json
//	@Param       id   path     string           true "id"
//	@Success     200  {object} responses.GetEvent
//	@Failure     400
//	@Failure     404
//	@Router      /events/{id} [get]
//	@Security    public
func (h *EventsHandler) GetEvent(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("EventsHandler: got GetEvent request")

	eventId := &requests.GetEvent{}

	err := eventId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("EventsHandler: parse request GetEvent: %v", eventId)

	res, err := h.events.GetEvent(context.Background(), eventId.ID)
	if err != nil {
		h.logger.Warnf("can't EventsService.GetEvent: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("EventsHandler: request GetEvent done")

	return handler.OkResponse(res)
}

// GetEventsByRange retrieves an array of events in the given time range
//
//	@Summary     Retrieve events items by range
//	@Description Given a start time and end time recieve events in that period
//	@Tags        public.events
//	@Accept      json
//	@Produce     json
//	@Param       request body requests.GetEventsByRange true "Range"
//	@Success     200  {object} responses.GetEventsByRange
//	@Failure     400
//	@Failure     404
//	@Router      /events/range/ [get]
//	@Security    public
func (h *EventsHandler) GetEventsByRange(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("EventsHandler: got GetEventsByRange request")

	timeRange := &requests.GetEventsByRange{}

	err := timeRange.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("EventsHandler: parse request GetEventsByRange: %v", timeRange)

	res, err := h.events.GetEventsByRange(context.Background(), timeRange.From, timeRange.To)
	if err != nil {
		h.logger.Warnf("can't EventsService.GetEventsByRange: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("EventsHandler: request GetEventsByRange done")

	return handler.OkResponse(res)
}

// PostEvent creates a new event item
//
//		@Summary     Create a new event item
//		@Description Create a new event item with the provided data
//		@Tags        auth.events
//		@Accept      json
//		@Param       request body requests.PostEvent true "Event data"
//		@Success     201
//		@Failure     400
//	 	@Failure     401
//		@Failure     404
//		@Router      /events [post]
//		@Security    Authorised
func (h *EventsHandler) PostEvent(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("EventsHandler: got PostEvent request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token PostEvent: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteEvent: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("EventsHandler: PostEvent Authenticated: %v", resp.MemberID)

	event := &requests.PostEvent{}

	err = event.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind PostEvent: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("EventsHandler: parse request PostEvent: %v", event)

	err = h.events.PostEvent(context.Background(), *mapper.MakeRequestPostEvent(*event))
	if err != nil {
		h.logger.Warnf("can't EventsService.PostEvent: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("EventsHandler: request PostEvent done")

	return handler.CreatedResponse(nil)
}

// DeleteEvent deletes an event item by ID
//
//	@Summary     Delete an event item by ID
//	@Description Delete a specific event item using its ID
//	@Tags        auth.events
//	@Param       id   path     string           true "Event ID"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Router      /events/{id} [delete]
//	@Security    Authorised
func (h *EventsHandler) DeleteEvent(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("EventsHandler: got DeleteEvent request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token DeleteEvent: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteEvent: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("EventsHandler: DeleteEvent Authenticated: %v", resp.MemberID)

	eventId := &requests.DeleteEvent{}

	err = eventId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind DeleteEvent: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("EventsHandler: parse request DeleteEvent: %v", eventId)

	err = h.events.DeleteEvent(context.Background(), eventId.ID)
	if err != nil {
		h.logger.Warnf("can't EventService.DeleteEvent: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("EventsHandler: request DeleteEvent done")

	return handler.OkResponse(nil)
}

// UpdateEvent updates a event item
//
//	@Summary     Update a event item
//	@Description Update an existing event item with the provided data
//	@Tags        auth.events
//	@Accept      json
//	@Param       id   path     string           true "Event ID"
//	@Param       request body requests.UpdateEvent true "Event new data"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Router      /events/{id} [put]
//	@Security    Authorised
func (h *EventsHandler) UpdateEvent(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("EventsHandler: got UpdateEvent request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token UpdateEvent: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteEvent: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("EventsHandler: UpdateEvent Authenticated: %v", resp.MemberID)

	event := &requests.UpdateEvent{}

	err = event.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind UpdateEvent: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("EventsHandler: parse request UpdateEvent: %v", event)

	err = h.events.UpdateEvent(context.Background(), *mapper.MakeRequestUpdateEvent(*event))
	if err != nil {
		h.logger.Warnf("can't EventService.UpdateEvent: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("EventsHandler: request UpdateEvent done")

	return handler.OkResponse(nil)
}
