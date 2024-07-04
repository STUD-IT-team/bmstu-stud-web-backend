package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

type ClubsHandler struct {
	r      handler.Renderer
	clubs  app.ClubService
	logger *log.Logger
	guard  *app.GuardService
}

func NewClubsHandler(r handler.Renderer, clubs app.ClubService, logger *log.Logger, guard *app.GuardService) *ClubsHandler {
	return &ClubsHandler{
		r:      r,
		clubs:  clubs,
		logger: logger,
		guard:  guard,
	}
}

func (h *ClubsHandler) BasePrefix() string {
	return "/clubs"
}

func (h *ClubsHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllClubs))
	r.Get("/{club_id}", h.r.Wrap(h.GetClub))
	r.Get("/type={type}", h.r.Wrap(h.GetClubsByType))
	r.Get("/search/{name}", h.r.Wrap(h.GetClubsByName))
	r.Get("/members/{club_id}", h.r.Wrap(h.GetClubMembers))
	r.Get("/media/{club_id}", h.r.Wrap(h.GetClubMedia))
	r.Post("/", h.r.Wrap(h.PostClub))
	r.Delete("/{club_id}", h.r.Wrap(h.DeleteClub))
	r.Put("/{club_id}", h.r.Wrap(h.UpdateClub))

	return r
}

func (h *ClubsHandler) GetAllClubs(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got getAllClub request")

	res, err := h.clubs.GetAllClubs()
	if err != nil {
		log.WithError(err).Warnf("can't service.GetAllClubs GetAllClubs")
		return handler.NotFoundResponse()
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

func (h *ClubsHandler) GetClub(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got getClub request")

	clubId := &requests.GetClub{}

	err := clubId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetClub: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubId)

	res, err := h.clubs.GetClub(clubId.ID)
	if err != nil {
		h.logger.Warnf("can't ClubService.GetClub: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

func (h *ClubsHandler) GetClubsByType(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubsByType request")

	clubName := &requests.GetClubsByType{}

	err := clubName.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetClubsByType: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubName)

	res, err := h.clubs.GetClubsByType(clubName.Type)
	if err != nil {
		h.logger.Warnf("can't ClubService.GetClub: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

func (h *ClubsHandler) GetClubsByName(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubByName request")

	clubName := &requests.GetClubsByName{}

	err := clubName.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetClubByName: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubName)

	res, err := h.clubs.GetClubsByName(clubName.Name)
	if err != nil {
		h.logger.Warnf("can't ClubService.GetClub: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

func (h *ClubsHandler) GetClubMembers(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubMembers request")

	clubId := &requests.GetClubMembers{}

	err := clubId.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetClubMembers GetClubMembers")
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubId)

	res, err := h.clubs.GetClubMembers(clubId.ID)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetClubMembers GetClubMembers")
		return handler.NotFoundResponse()
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

func (h *ClubsHandler) GetClubMedia(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubMedia request")
	clubId := &requests.GetClubMedia{}

	err := clubId.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetClubMedia GetClubMedia")
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubId)

	res, err := h.clubs.GetClubMediaFiles(clubId.ID)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetClubMedia GetClubMedia")
		return handler.NotFoundResponse()
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

func (h *ClubsHandler) PostClub(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got PostClub request")

	access, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: access})
	if err != nil || !resp.Valid {
		h.logger.Warnf("Unauthorized request: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("ClubsHandler: PostClub Authenticated: %v", resp.MemberID)

	club := &requests.PostClub{}
	err = club.Bind(req)
	if err != nil {
		h.logger.Warnf("can't service.PostClub %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", club)

	err = h.clubs.PostClub(context.Background(), club)
	if err != nil {
		h.logger.Warnf("can't service.PostClub %v", err)
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("ClubsHandler: request done")
	return handler.OkResponse(nil)
}

func (h *ClubsHandler) DeleteClub(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got DeleteClub request")

	access, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: access})
	if err != nil || !resp.Valid {
		h.logger.Warnf("Unauthorized request: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("ClubsHandler: DeleteClub Authenticated: %v", resp.MemberID)

	club := &requests.DeleteClub{}

	err = club.Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.DeleteClub DeleteClub")
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: Parsed request: %v", club)

	err = h.clubs.DeleteClub(club.ID)
	if err != nil {
		log.WithError(err).Warnf("can't service.DeleteClub DeleteClub")
		return handler.InternalServerErrorResponse()
	}

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
