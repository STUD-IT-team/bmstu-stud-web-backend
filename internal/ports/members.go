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
	r.Get("/{id}", h.r.Wrap(h.GetMember))
	r.Get("/search/{name}", h.r.Wrap(h.GetMembersByName))
	r.Post("/", h.r.Wrap(h.PostMember))
	r.Delete("/{id}", h.r.Wrap(h.DeleteMember))
	r.Put("/{id}", h.r.Wrap(h.UpdateMember))

	return r
}

func (h *MembersHandler) GetAllMembers(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetAllMembers request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: GetAllMembers Authenticated: %v", resp.MemberID)

	res, err := h.members.GetAllMembers(context.Background())
	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetAllMembers: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("MembersHandler: request GetAllMembers done")

	return handler.OkResponse(res)
}

func (h *MembersHandler) GetMember(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetMember request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: GetAllMembers Authenticated: %v", resp.MemberID)

	memberId := &requests.GetMember{}

	err = memberId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MembersHandler: parse request GetMember: %v", memberId)

	res, err := h.members.GetMember(context.Background(), memberId.ID)
	if err != nil {
		h.logger.Warnf("can't MembersService.GetMember: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("MembersHandler: request GetMember done")

	return handler.OkResponse(res)
}

func (h *MembersHandler) GetMembersByName(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetMembersByName request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: GetAllMembers Authenticated: %v", resp.MemberID)

	name := &requests.GetMembersByName{}

	err = name.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind GetMembersByName: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MembersHandler: parse request GetMembersByName: %v", name)

	res, err := h.members.GetMembersByName(context.Background(), name.Search)
	if err != nil {
		h.logger.Warnf("can't MemberService.GetMembersByName: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("MembersHandler: request GetMembersByName done")

	return handler.OkResponse(res)
}

func (h *MembersHandler) PostMember(w http.ResponseWriter, req *http.Request) handler.Response {
	// member := &requests.PostMember{}

	// err := member.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostMember PostMember")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.members.PostMembers(context.Background(), member)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.PostMember PostMember")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *MembersHandler) DeleteMember(w http.ResponseWriter, req *http.Request) handler.Response {
	// member := &requests.DeleteMember{}

	// err := member.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteMember DeleteMember")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.members.DeleteMember(context.Background(), member.ID)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.DeleteMember DeleteMember")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}

func (h *MembersHandler) UpdateMember(w http.ResponseWriter, req *http.Request) handler.Response {
	// member := &requests.UpdateMember{}

	// err := member.Bind(req)
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateMember UpdateMember")
	// 	return handler.BadRequestResponse()
	// }

	// err = h.members.UpdateMember(context.Background(), *mapper.MakeRequestPutMember(*member))
	// if err != nil {
	// 	log.WithError(err).Warnf("can't service.UpdateMember UpdateMember")
	// 	return handler.InternalServerErrorResponse()
	// }

	return handler.OkResponse(nil)
}
