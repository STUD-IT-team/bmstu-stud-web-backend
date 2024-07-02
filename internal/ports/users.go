package http

import (
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
)

type UsersHandler struct {
	r     handler.Renderer
	users app.UsersService
}

func NewUsersHandler(r handler.Renderer, users app.UsersService) *UsersHandler {
	return &UsersHandler{
		r:     r,
		users: users,
	}
}

func (h *UsersHandler) BasePrefix() string {
	return "/users"
}

func (h *UsersHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllUsers))
	r.Get("/{id}", h.r.Wrap(h.GetUser))
	r.Get("/{filters}", h.r.Wrap(h.GetUsers))
	r.Post("/{id}", h.r.Wrap(h.PostUser))
	r.Delete("/{id}", h.r.Wrap(h.DeleteUser))
	r.Put("/{id}", h.r.Wrap(h.UpdateUser))

	return r
}

func (h *UsersHandler) GetAllUsers(w http.ResponseWriter, req *http.Request) handler.Response {
	// filter := &requests.GetAllUsers{}

	// res, err := h.users.GetAllUsers(context.Background(), *filter)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetAllUsers GetAllUsers")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *UsersHandler) GetUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// userId := &requests.GetUser{}

	// err := userId.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetUser GetUser")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.users.GetUser(context.Background(), userId.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetUser GetUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *UsersHandler) GetUsers(w http.ResponseWriter, req *http.Request) handler.Response {
	// filter := &requests.GetUsers{}

	// err := filter.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetUsers GetUsers")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.users.GetUsers(context.Background(), filter)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetUsers GetUsers")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *UsersHandler) PostUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// user := &requests.PostUser{}

	// err := user.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostUser PostUser")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.users.PostUsers(context.Background(), user)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostUser PostUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *UsersHandler) DeleteUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// user := &requests.DeleteUser{}

	// err := user.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteUser DeleteUser")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.users.DeleteUser(context.Background(), user.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteUser DeleteUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *UsersHandler) UpdateUser(w http.ResponseWriter, req *http.Request) handler.Response {
	// user := &requests.UpdateUser{}

	// err := user.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateUser UpdateUser")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.users.UpdateUser(context.Background(), *mapper.MakeRequestPutUser(*user))
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateUser UpdateUser")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}
