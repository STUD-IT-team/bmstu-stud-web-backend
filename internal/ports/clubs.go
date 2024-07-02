package http

import (
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

type ClubsHandler struct {
	r     handler.Renderer
	clubs app.ClubService
}

func NewClubsHandler(r handler.Renderer, clubs app.ClubService) *ClubsHandler {
	return &ClubsHandler{
		r:     r,
		clubs: clubs,
	}
}

func (h *ClubsHandler) BasePrefix() string {
	return "/clubs"
}

func (h *ClubsHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllClubs))
	r.Get("/{club_id}", h.r.Wrap(h.GetClub))
	r.Get("/{type}", h.r.Wrap(h.GetClubsByType))
	r.Get("/search/{name}", h.r.Wrap(h.GetClubsByName))
	r.Get("/members/{club_id}", h.r.Wrap(h.GetClubMembers))
	r.Get("/media/{club_id}", h.r.Wrap(h.GetClubMedia))
	r.Post("/{club_id}", h.r.Wrap(h.PostClub))
	r.Delete("/{club_id}", h.r.Wrap(h.DeleteClub))
	r.Put("/{club_id}", h.r.Wrap(h.UpdateClub))

	return r
}

func (h *ClubsHandler) GetAllClubs(w http.ResponseWriter, req *http.Request) handler.Response {
	// filter := &requests.GetAllClubs{}

	// res, err := h.clubs.GetClubs(context.Background(), *filter)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetAllClubs GetAllClubs")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *ClubsHandler) GetClub(w http.ResponseWriter, req *http.Request) handler.Response {
	clubId := &requests.GetClub{}

	err := clubId.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't requests.Bind GetClub: %v", err)
		return handler.BadRequestResponse()
	}

	res, err := h.clubs.GetClub(clubId.ID)
	if err != nil {
		log.WithError(err).Warnf("can't ClubService.GetClub: %v", err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *ClubsHandler) GetClubsByType(w http.ResponseWriter, req *http.Request) handler.Response {
	// filter := &requests.GetClubsByType{}

	// err := filter.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetClubsByType GetClubsByType")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.clubs.GetClubsByFilter(context.Background(), *filter)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetClubsByType GetClubsByType")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *ClubsHandler) GetClubsByName(w http.ResponseWriter, req *http.Request) handler.Response {
	// filter := &requests.GetClubsByName{}

	// err := filter.Bind(req)
	// res, err := h.clubs.GetClubsByFilter(context.Background(), *filter)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetClubsByName GetClubsByName")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *ClubsHandler) GetClubMembers(w http.ResponseWriter, req *http.Request) handler.Response {
	// clubId := &requests.GetClubMembers{}

	// err := clubId.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetClubMembers GetClubMembers")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.clubs.GetClubMembers(context.Background(), clubId.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetClubMembers GetClubMembers")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *ClubsHandler) GetClubMedia(w http.ResponseWriter, req *http.Request) handler.Response {
	// clubId := &requests.GetClubMedia{}

	// err := clubId.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetClubMedia GetClubMedia")
	// 	return handler.BadRequestResponse()
	// }

	// res, err := h.clubs.GetClubMedia(context.Background(), clubId.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.GetClubMedia GetClubMedia")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *ClubsHandler) PostClub(w http.ResponseWriter, req *http.Request) handler.Response {
	// club := &requests.PostClub{}

	// err := club.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostClub PostClub")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.clubs.PostClub(context.Background(), club)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostClub PostClub")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *ClubsHandler) DeleteClub(w http.ResponseWriter, req *http.Request) handler.Response {
	// club := &requests.DeleteClub{}

	// err := club.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteClub DeleteClub")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.clubs.DeleteClub(context.Background(), club.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteClub DeleteClub")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *ClubsHandler) UpdateClub(w http.ResponseWriter, req *http.Request) handler.Response {
	// club := &requests.UpdateClub{}

	// err := club.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateClub UpdateClub")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.clubs.UpdateClub(context.Background(), club)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateClub UpdateClub")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}
