package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
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
	r.Get("/type/{type}", h.r.Wrap(h.GetClubsByType))
	r.Get("/search/{name}", h.r.Wrap(h.GetClubsByName))
	r.Get("/members/{club_id}", h.r.Wrap(h.GetClubMembers))
	r.Get("/media/{club_id}", h.r.Wrap(h.GetClubMedia))
	r.Post("/", h.r.Wrap(h.PostClub))
	r.Delete("/{club_id}", h.r.Wrap(h.DeleteClub))
	r.Put("/{club_id}", h.r.Wrap(h.UpdateClub))
	r.Post("/media/{club_id}", h.r.Wrap(h.PostClubMedia))
	r.Delete("/media/{club_id}", h.r.Wrap(h.DeleteClubMedia))
	r.Put("/media/{club_id}", h.r.Wrap(h.UpdateClubMedia))
	r.Get("/clearance/post/", h.r.Wrap(h.GetClearancePost))
	r.Get("/clearance/delete/{club_id}", h.r.Wrap(h.GetClearanceDelete))
	r.Get("/clearance/update/{club_id}", h.r.Wrap(h.GetClearanceUpdate))
	r.Get("/media/clearance/post/{club_id}", h.r.Wrap(h.GetMediaClearancePost))
	r.Get("/media/clearance/delete/{club_id}", h.r.Wrap(h.GetMediaClearanceDelete))
	r.Get("/media/clearance/update/{club_id}", h.r.Wrap(h.GetMediaClearanceUpdate))

	return r
}

func (h *ClubsHandler) GetClearancePost(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClearancePost request")

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

	h.logger.Infof("ClubsHandler: GetClearancePost Authenticated: %v", resp.MemberID)

	response, err := h.clubs.GetClearancePost(context.Background(), resp)
	if err != nil {
		h.logger.Warnf("can't clubs.GetClearancePost: %v", err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(response)
}

func (h *ClubsHandler) GetClearanceDelete(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClearanceDelete request")

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

	h.logger.Infof("ClubsHandler: GetClearanceDelete Authenticated: %v", resp.MemberID)

	response, err := h.clubs.GetClearanceDelete(context.Background(), resp)
	if err != nil {
		h.logger.Warnf("can't clubs.GetClearanceDelete: %v", err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(response)
}

func (h *ClubsHandler) GetClearanceUpdate(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClearanceUpdate request")

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

	h.logger.Infof("ClubsHandler: GetClearanceUpdate Authenticated: %v", resp.MemberID)

	parsed_req := &requests.GetClearanceClubUpdate{}
	err = parsed_req.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request: %v", err)
		return handler.BadRequestResponse()
	}

	response, err := h.clubs.GetClearanceUpdate(context.Background(), resp, parsed_req.ClubID)
	if err != nil {
		h.logger.Warnf("can't clubs.GetClearanceUpdate: %v", err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(response)
}

func (h *ClubsHandler) GetMediaClearancePost(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetMediaClearancePost request")

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

	h.logger.Infof("ClubsHandler: GetMediaClearancePost Authenticated: %v", resp.MemberID)

	parsed_req := &requests.GetClearanceClubUpdate{}
	err = parsed_req.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request: %v", err)
		return handler.BadRequestResponse()
	}

	response, err := h.clubs.GetClearanceMediaPost(context.Background(), resp, parsed_req.ClubID)
	if err != nil {
		h.logger.Warnf("can't clubs.GetMediaClearancePost: %v", err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(response)
}

func (h *ClubsHandler) GetMediaClearanceDelete(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetMediaClearanceDelete request")

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

	h.logger.Infof("ClubsHandler: GetMediaClearanceDelete Authenticated: %v", resp.MemberID)

	parsed_req := &requests.GetClearanceClubUpdate{}
	err = parsed_req.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request: %v", err)
		return handler.BadRequestResponse()
	}

	response, err := h.clubs.GetClearanceMediaDelete(context.Background(), resp, parsed_req.ClubID)
	if err != nil {
		h.logger.Warnf("can't clubs.GetMediaClearanceDelete: %v", err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(response)
}

func (h *ClubsHandler) GetMediaClearanceUpdate(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetMediaClearanceUpdate request")

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

	h.logger.Infof("ClubsHandler: GetMediaClearanceUpdate Authenticated: %v", resp.MemberID)

	parsed_req := &requests.GetClearanceClubUpdate{}
	err = parsed_req.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse request: %v", err)
		return handler.BadRequestResponse()
	}

	response, err := h.clubs.GetClearanceMediaUpdate(context.Background(), resp, parsed_req.ClubID)
	if err != nil {
		h.logger.Warnf("can't clubs.GetMediaClearanceUpdate: %v", err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(response)
}

// GetAllClubs
//
// @Summary    Возвращает все клубы из БД
// @Description  Возвращает все клубы для страницы с поиском клубов
// @Tags      public.club
// @Produce    json
// @Success    200      {object}  responses.GetAllClubs
// @Failure    404
// @Failure    500
// @Router      /clubs [get]
// @Security    Public
func (h *ClubsHandler) GetAllClubs(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got getAllClub request")

	res, err := h.clubs.GetAllClubs(context.Background())
	if err != nil {
		h.logger.Warnf("can't service.GetAllClubs GetAllClubs: %v", err)
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

// GetClub
//
// @Summary    Возвращает клуб по ID
// @Description  Возвращает информацию о клубе для страницы клуба
// @Tags      public.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Success    200      {object}  responses.GetClub
// @Failure    400
// @Failure    404
// @Failure    500
// @Router      /clubs/{club_id} [get]
// @Security    Public
func (h *ClubsHandler) GetClub(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got getClub request")

	clubId := &requests.GetClub{}

	err := clubId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetClub: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubId)

	res, err := h.clubs.GetClub(context.Background(), clubId.ID)
	if err != nil {
		h.logger.Warnf("can't ClubService.GetClub: %v", err)
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

// GetClubsByType
//
// @Summary    Возвращает все клубы по типу
// @Description  Возвращает все клубы, у которых введенная строка является подстрокой в типе
// @Tags      public.club
// @Produce    json
// @Param      club_type    path    string  true  "club type"
// @Success    200      {object}  responses.GetClubsByType
// @Failure    400
// @Failure    404
// @Failure    500
// @Router      /clubs/type/{club_type} [get]
// @Security    Public
func (h *ClubsHandler) GetClubsByType(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubsByType request")

	clubName := &requests.GetClubsByType{}

	err := clubName.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetClubsByType: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubName)

	res, err := h.clubs.GetClubsByType(context.Background(), clubName.Type)
	if err != nil {
		h.logger.Warnf("can't ClubService.GetClub: %v", err)
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

// GetClubsByName
//
// @Summary    Возвращает все клубы по имени
// @Description  Возвращает все клубы, у которых введенная строка является подстрокой в имени
// @Tags      public.club
// @Produce    json
// @Param      club_name    path    string  true  "club name"
// @Success    200      {object}  responses.GetClubsByName
// @Failure    400
// @Failure    404
// @Failure    500
// @Router      /clubs/search/{club_name} [get]
// @Security    Public
func (h *ClubsHandler) GetClubsByName(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubByName request")

	clubName := &requests.GetClubsByName{}

	err := clubName.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetClubByName: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubName)

	res, err := h.clubs.GetClubsByName(context.Background(), clubName.Name)
	if err != nil {
		h.logger.Warnf("can't ClubService.GetClub: %v", err)
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

// GetClubMembers
//
// @Summary    Возвращает руководителя клуба и руководителей подклубов
// @Description  Возвращает руководителя клуба и руководителей подклубов(1-го уровня) по ID клуба
// @Tags      public.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Success    200      {object}  responses.GetClubMembers
// @Failure    400
// @Failure    404
// @Failure    500
// @Router      /clubs/members/{club_id} [get]
// @Security    Public
func (h *ClubsHandler) GetClubMembers(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubMembers request")

	clubId := &requests.GetClubMembers{}

	err := clubId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't service.GetClubMembers GetClubMembers: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubId)

	res, err := h.clubs.GetClubMembers(context.Background(), clubId.ID)
	if err != nil {
		h.logger.Warnf("can't service.GetClubMembers GetClubMembers: %v", err)
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

// GetClubMedia
//
// @Summary    Возвращает фотографии клуба
// @Description  Возвращает фотографии клуба по его ID
// @Tags      public.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Success    200      {object}  responses.GetClubMedia
// @Failure    400
// @Failure    404
// @Failure    500
// @Router      /clubs/media/{club_id} [get]
// @Security    Public
func (h *ClubsHandler) GetClubMedia(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got GetClubMedia request")
	clubId := &requests.GetClubMedia{}

	err := clubId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't service.GetClubMedia GetClubMedia: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request: %v", clubId)

	res, err := h.clubs.GetClubMediaFiles(context.Background(), clubId.ID)
	if err != nil {
		h.logger.Warnf("can't service.GetClubMedia GetClubMedia: %v", err)
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(res)
}

// PostClub
//
// @Summary    Добавляет клуб в базу данных
// @Description Добавляет клуб в базу данных, требуется аутентификация
// @Tags      auth.club
// @Produce    json
// @Param      request  body    requests.PostClub  true  "post club data"
// @Success    200      {object}  responses.PostClubResponse
// @Failure    400
// @Failure    401
// @Failure    403
// @Failure    409
// @Failure    500
// @Router      /clubs [post]
// @Security    Authorized
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

	permission, err := h.clubs.GetClearancePost(context.Background(), resp)
	if err != nil {
		h.logger.Warnf("can't service.GetClearancePost %v", err)
		return handler.InternalServerErrorResponse()
	}
	if !permission.Access {
		h.logger.Warnf("do not have enough access for PostClub: %v", permission.Comment)
		return handler.ForbiddenResponse()
	}

	clubID, err := h.clubs.PostClub(context.Background(), club)

	if err != nil {
		h.logger.Warnf("can't service.PostClub %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")
	return handler.OkResponse(responses.PostClubResponse{ID: clubID})
}

// DeleteClub
//
// @Summary    Удаляет клуб из базу данных
// @Description Удаляет клуб в базу данных, а также все объеуты ассоциируемые с ним, требуется аутентификация
// @Tags      auth.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Success    200
// @Failure    400
// @Failure    401
// @Failure    404
// @Failure    500
// @Router      /clubs [delete]
// @Security    Authorized
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
		h.logger.Warnf("can't service.DeleteClub DeleteClub: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: Parsed request: %v", club)

	permission, err := h.clubs.GetClearanceDelete(context.Background(), resp)
	if err != nil {
		h.logger.Warnf("can't service.GetClearanceDelete %v", err)
		return handler.InternalServerErrorResponse()
	}
	if !permission.Access {
		h.logger.Warnf("do not have enough access for DeleteClub: %v", permission.Comment)
		return handler.ForbiddenResponse()
	}

	err = h.clubs.DeleteClub(context.Background(), club.ID)
	if err != nil {
		h.logger.Warnf("can't service.DeleteClub DeleteClub: %v", err)
		if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}
	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(nil)
}

// UpdateClub
//
// @Summary    Обновляет данные клуба в базе данных
// @Description Обновляет данные клуба в базе данных, требуется аутентификация
// @Tags      auth.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Param      request  body    requests.PostClub  true  "update club data"
// @Success    200
// @Failure    400
// @Failure    401
// @Failure    409
// @Failure    404
// @Failure    500
// @Router      /clubs [put]
// @Security    Authorized
func (h *ClubsHandler) UpdateClub(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got UpdateClub request")

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

	h.logger.Infof("ClubsHandler: UpdateClub Authenticated: %v", resp.MemberID)

	club := &requests.UpdateClub{}

	err = club.Bind(req)
	if err != nil {
		h.logger.Warnf("can't service.UpdateClub UpdateClub: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: Parsed request: %v", club)

	permission, err := h.clubs.GetClearanceUpdate(context.Background(), resp, club.ID)
	if err != nil {
		h.logger.Warnf("can't service.GetClearanceUpdate %v", err)
		return handler.InternalServerErrorResponse()
	}
	if !permission.Access {
		h.logger.Warnf("do not have enough access for UpdateClub: %v", permission.Comment)
		return handler.ForbiddenResponse()
	}

	err = h.clubs.UpdateClub(context.Background(), club)
	if err != nil {
		h.logger.Warnf("can't service.UpdateClub UpdateClub: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("ClubsHandler: request done")

	return handler.OkResponse(nil)
}

// PostClubPhoto
//
// @Summary    Добавляет в клуб фотографии клуба в базу данных
// @Description Добавляет в клуб фотографии клуба в базу данных
// @Tags      auth.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Param      request  body    requests.PostClubPhoto  true  "post club photo data"
// @Success    200
// @Failure    400
// @Failure    401
// @Failure    409
// @Failure    500
// @Router      /clubs/media/{club_id} [post]
// @Security    Authorized
func (h *ClubsHandler) PostClubMedia(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got PostClubMedia request")

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

	h.logger.Infof("ClubsHandler: PostClubMedia Authenticated: %v", resp.MemberID)

	photo := &requests.PostClubPhoto{}
	err = photo.Bind(req)
	if err != nil {
		h.logger.Warnf("can't service.PostClubMedia %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("ClubsHandler: parse request.")

	permission, err := h.clubs.GetClearanceMediaPost(context.Background(), resp, photo.ClubID)
	if err != nil {
		h.logger.Warnf("can't service.GetClearanceMediaPost %v", err)
		return handler.InternalServerErrorResponse()
	}
	if !permission.Access {
		h.logger.Warnf("do not have enough access for PostClubMedia: %v", permission.Comment)
		return handler.ForbiddenResponse()
	}

	err = h.clubs.PostClubPhoto(context.Background(), photo)
	if err != nil {
		h.logger.Warnf("can't service.PostClubMedia %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	return handler.CreatedResponse(nil)
}

// DeleteClubPhoto
//
// @Summary    Удаляет фотографию из фотографий клуба
// @Description Удаляет фотографию из фотографий клуба
// @Tags      auth.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Param      request  body    requests.DeleteClubPhoto  true  "post club photo data"
// @Success    200
// @Failure    400
// @Failure    401
// @Failure    404
// @Failure    500
// @Router      /clubs/media/{club_id} [delete]
// @Security    Authorized
func (h *ClubsHandler) DeleteClubMedia(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got DeleteClubMedia request")

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

	h.logger.Infof("ClubsHandler: DeleteClubMedia Authenticated: %v", resp.MemberID)

	photo := &requests.DeleteClubPhoto{}
	err = photo.Bind(req)
	if err != nil {
		h.logger.Warnf("can't parse DeleteClubMedia %v", err)
		return handler.BadRequestResponse()
	}
	h.logger.Infof("ClubsHandler: parse request.")

	permission, err := h.clubs.GetClearanceMediaDelete(context.Background(), resp, photo.ClubID)
	if err != nil {
		h.logger.Warnf("can't service.GetClearanceMediaDelete %v", err)
		return handler.InternalServerErrorResponse()
	}
	if !permission.Access {
		h.logger.Warnf("do not have enough access for DeleteClubMedia: %v", permission.Comment)
		return handler.ForbiddenResponse()
	}

	err = h.clubs.DeleteClubPhoto(context.Background(), photo)
	if err != nil {
		h.logger.Warnf("can't service.DeleteClubMedia %v", err)
		if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(nil)
}

// UpdateClubPhoto
//
// @Summary    Обновляет все фотографии клуба
// @Description Обновляет все фотографии клуба
// @Tags      auth.club
// @Produce    json
// @Param      club_id    path    int  true  "club id"
// @Param      request  body    requests.UpdateClubPhoto  true  "post club photo data"
// @Success    200
// @Failure    400
// @Failure    401
// @Failure    409
// @Failure    500
// @Router      /clubs/media/{club_id} [put]
// @Security    Authorized
func (h *ClubsHandler) UpdateClubMedia(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("ClubsHandler: got UpdateClubMedia request")

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

	h.logger.Infof("ClubsHandler: UpdateClubMedia Authenticated: %v", resp.MemberID)

	photo := &requests.UpdateClubPhoto{}

	err = photo.Bind(req)

	if err != nil {
		h.logger.Warnf("can't parse UpdateClubMedia %v", err)
		return handler.BadRequestResponse()
	}
	h.logger.Infof("ClubsHandler: parse request.")

	permission, err := h.clubs.GetClearanceMediaUpdate(context.Background(), resp, photo.ClubID)
	if err != nil {
		h.logger.Warnf("can't service.GetClearanceMediaUpdate %v", err)
		return handler.InternalServerErrorResponse()
	}
	if !permission.Access {
		h.logger.Warnf("do not have enough access for UpdateClubMedia: %v", permission.Comment)
		return handler.ForbiddenResponse()
	}

	err = h.clubs.UpdateClubPhoto(context.Background(), photo)
	if err != nil {
		h.logger.Warnf("can't service.UpdateClubMedia %v", err)
		if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		}
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(nil)
}
