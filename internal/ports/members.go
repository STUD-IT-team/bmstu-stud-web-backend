package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

type MembersHandler struct {
	r       handler.Renderer
	members app.MembersService
	logger  *log.Logger
	guard   *app.GuardService
}

func NewMembersHandler(r handler.Renderer, members app.MembersService, logger *log.Logger, guard *app.GuardService) *MembersHandler {
	return &MembersHandler{
		r:       r,
		members: members,
		logger:  logger,
		guard:   guard,
	}
}

func (h *MembersHandler) BasePrefix() string {
	return "/members"
}

func (h *MembersHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllMembers))
	r.Get("/{id}", h.r.Wrap(h.GetUser))
	r.Get("/search/", h.r.Wrap(h.GetMembersByFilter))
	r.Post("/", h.r.Wrap(h.PostUser))
	r.Delete("/{id}", h.r.Wrap(h.DeleteUser))
	r.Put("/{id}", h.r.Wrap(h.UpdateUser))

	return r
}

func (h *MembersHandler) GetAllMembers(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetAllMembers request")

	res, err := h.members.GetAllMembers(context.Background())
	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetAllMembers: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("MembersHandler: request GetAllMembers done")

	return handler.OkResponse(res)
}

func (h *MembersHandler) GetUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// userId := &requests.GetUser{}

	// err := userId.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetUser GetUser")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.members.GetUser(context.Background(), userId.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetUser GetUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *MembersHandler) GetMembersByFilter(w http.ResponseWriter, req *http.Request) handler.Response {
	// filter := &requests.GetMembers{}

	// err := filter.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetMembers GetMembers")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.members.GetMembers(context.Background(), filter)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetMembers GetMembers")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *MembersHandler) PostUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// user := &requests.PostUser{}

	// err := user.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostUser PostUser")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.members.PostMembers(context.Background(), user)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostUser PostUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *MembersHandler) DeleteUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// user := &requests.DeleteUser{}

	// err := user.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteUser DeleteUser")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.members.DeleteUser(context.Background(), user.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteUser DeleteUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *MembersHandler) UpdateUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// user := &requests.UpdateUser{}

	// err := user.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateUser UpdateUser")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.members.UpdateUser(context.Background(), *mapper.MakeRequestPutUser(*user))
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateUser UpdateUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}
