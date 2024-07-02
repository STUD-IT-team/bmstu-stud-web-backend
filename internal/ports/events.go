package http

// import (
// 	"context"
// 	"net/http"

// 	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
// 	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
// 	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
// 	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

// 	"github.com/go-chi/chi"
// 	log "github.com/sirupsen/logrus"
// )

// type EventsHandler struct {
// 	r      handler.Renderer
// 	events app.EventsService
// }

// func NewEventsHandler(r handler.Renderer, events app.EventsService) *EventsHandler {
// 	return &EventsHandler{
// 		r:      r,
// 		events: events,
// 	}
// }

// func (h *EventsHandler) BasePrefix() string {
// 	return "/events"
// }

// func (h *EventsHandler) Routes() chi.Router {
// 	r := chi.NewRouter()

// 	r.Get("/", h.r.Wrap(h.GetAllEvents))
// 	r.Get("/{id}", h.r.Wrap(h.GetEvent))
// 	r.Post("/{event}", h.r.Wrap(h.PostEvent))
// 	r.Delete("/{id}", h.r.Wrap(h.DeleteEvent))
// 	r.Put("/{id}", h.r.Wrap(h.UpdateEvent))

// 	return r
// }

// func (h *EventsHandler) GetAllEvents(w http.ResponseWriter, req *http.Request) handler.Response {
// 	filter := &requests.GetAllEvents{}

// 	res, err := h.events.GetAllEvents(context.Background(), *filter)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.GetAllEvents GetAllEvents")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(res)
// }

// func (h *EventsHandler) GetEvent(w http.ResponseWriter, req *http.Request) handler.Response {
// 	eventId := &requests.GetEvent{}

// 	err := eventId.Bind(req)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.GetEvent GetEvent")
// 		return handler.BadRequestResponse()
// 	}

// 	res, err := h.events.GetEvent(context.Background(), eventId.ID)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.GetEvent GetEvent")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(res)
// }

// func (h *EventsHandler) PostEvent(w http.ResponseWriter, req *http.Request) handler.Response {
// 	event := &requests.PostEvent{}

// 	err := event.Bind(req)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.PostEvent PostEvent")
// 		return handler.BadRequestResponse()
// 	}

// 	err = h.events.PostEvents(context.Background(), *mapper.MakeRequestPutEvents(*event))
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.PostEvent PostEvent")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(nil)
// }

// func (h *EventsHandler) DeleteEvent(w http.ResponseWriter, req *http.Request) handler.Response {
// 	event := &requests.DeleteEvent{}

// 	err := event.Bind(req)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.DeleteEvent DeleteEvent")
// 		return handler.BadRequestResponse()
// 	}

// 	err = h.events.DeleteEvent(context.Background(), event.ID)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.DeleteEvent DeleteEvent")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(nil)
// }

// func (h *EventsHandler) UpdateEvent(w http.ResponseWriter, req *http.Request) handler.Response {
// 	event := &requests.UpdateEvent{}

// 	err := event.Bind(req)
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.UpdateEvent UpdateEvent")
// 		return handler.BadRequestResponse()
// 	}

// 	err = h.events.UpdateEvent(context.Background(), *mapper.MakeRequestPutEvent(*event))
// 	if err != nil {
// 		log.WithError(err).Warnf("can't service.UpdateEvent UpdateEvent")
// 		return handler.InternalServerErrorResponse()
// 	}

// 	return handler.OkResponse(nil)
// }
