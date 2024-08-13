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
	r.Get("/clearance/", h.r.Wrap(h.GetClearance))

	return r
}

// GetAllMembers retrieves all members
//
//	@Summary     Retrieve all members
//	@Description Get a list of all members
//	@Tags        auth.members
//	@Produce     json
//	@Success     200 {object} responses.GetAllMembers
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Failure     500
//	@Router      /members/ [get]
//	@Security    Authorised
func (h *MembersHandler) GetAllMembers(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetAllMembers request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	checkResp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !checkResp.Valid {
		h.logger.Warnf("can't GuardService.Check on GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: GetAllMembers Authenticated: %v", checkResp.MemberID)

	cleaResp, err := h.members.GetClearance(context.Background(), checkResp)

	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetClearance: %v", err)
		return handler.InternalServerErrorResponse()
	}

	if !cleaResp.Access {
		h.logger.Warnf("Not allowed: %s", cleaResp.Comment)
		return handler.ForbiddenResponse()
	}

	h.logger.Infof("MembersHandler: GetAllMembers Allowed: %v", checkResp.MemberID)

	res, err := h.members.GetAllMembers(context.Background())
	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetAllMembers: %v", err)
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("MembersHandler: request GetAllMembers done")

	return handler.OkResponse(res)
}

// GetMember retrieves a member by ID
//
//	@Summary     Retrieve member by ID
//	@Description Get a specific member using its ID
//	@Tags        auth.members
//	@Produce     json
//	@Param       id   path     string           true "Member ID"
//	@Success     200 {object} responses.GetMember
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Failure     500
//	@Router      /members/{id} [get]
//	@Security    Authorised
func (h *MembersHandler) GetMember(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetMember request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	checkResp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !checkResp.Valid {
		h.logger.Warnf("can't GuardService.Check on GetAllMembers: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: GetAllMembers Authenticated: %v", checkResp.MemberID)

	cleaResp, err := h.members.GetClearance(context.Background(), checkResp)

	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetClearance: %v", err)
		return handler.InternalServerErrorResponse()
	}

	if !cleaResp.Access {
		h.logger.Warnf("Not allowed: %s", cleaResp.Comment)
		return handler.ForbiddenResponse()
	}

	h.logger.Infof("MembersHandler: GetMember Allowed: %v", checkResp.MemberID)

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
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("MembersHandler: request GetMember done")

	return handler.OkResponse(res)
}

// GetMembersByName retrieves members by name
//
//	@Summary     Retrieve members by name
//	@Description Get members that match the specified name
//	@Tags        auth.members
//	@Produce     json
//	@Param       name   path     string           true "Member name"
//	@Success     200 {array} responses.GetMembersByName
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Failure     500
//	@Router      /members/search/{name} [get]
//	@Security    Authorised
func (h *MembersHandler) GetMembersByName(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetMembersByName request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token GetMembersByName: %v", err)
		return handler.UnauthorizedResponse()
	}

	checkResp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !checkResp.Valid {
		h.logger.Warnf("can't GuardService.Check on GetMembersByName: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: GetMembersByName Authenticated: %v", checkResp.MemberID)

	cleaResp, err := h.members.GetClearance(context.Background(), checkResp)

	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetClearance: %v", err)
		return handler.InternalServerErrorResponse()
	}

	if !cleaResp.Access {
		h.logger.Warnf("Not allowed: %s", cleaResp.Comment)
		return handler.ForbiddenResponse()
	}

	h.logger.Infof("MembersHandler: GetMembersByName Allowed: %v", checkResp.MemberID)

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
		if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("MembersHandler: request GetMembersByName done")

	return handler.OkResponse(res)
}

// PostMember creates a new member
//
//	@Summary     Create a new member
//	@Description Create a new member with the provided data
//	@Tags        auth.members
//	@Accept      json
//	@Param       request body requests.PostMember true "Member data"
//	@Success     201
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Failure     409
//	@Failure     500
//	@Router      /members/ [post]
//	@Security    Authorised
func (h *MembersHandler) PostMember(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got PostMember request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token PostMember: %v", err)
		return handler.UnauthorizedResponse()
	}

	checkResp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !checkResp.Valid {
		h.logger.Warnf("can't GuardService.Check on PostMember: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: PostMember Authenticated: %v", checkResp.MemberID)

	cleaResp, err := h.members.GetClearance(context.Background(), checkResp)

	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetClearance: %v", err)
		return handler.InternalServerErrorResponse()
	}

	if !cleaResp.Access {
		h.logger.Warnf("Not allowed: %s", cleaResp.Comment)
		return handler.ForbiddenResponse()
	}

	h.logger.Infof("MembersHandler: PostMember Allowed: %v", checkResp.MemberID)

	member := &requests.PostMember{}

	err = member.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind PostMember: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MembersHandler: parse request PostMember: %v", member)

	err = h.members.PostMember(context.Background(), mapper.MakeRequestPostMember(member))
	if err != nil {
		h.logger.Warnf("can't MembersService.PostMember: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("MembersHandler: request PostMember done")

	return handler.CreatedResponse(nil)
}

// DeleteMember deletes a member by ID
//
//	@Summary     Delete a member by ID
//	@Description Delete a member using its ID
//	@Tags        auth.members
//	@Produce     json
//	@Param       id   path     string           true "Member ID"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Failure     500
//	@Router      /members/{id} [delete]
//	@Security    Authorised
func (h *MembersHandler) DeleteMember(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got DeleteMember request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token DeleteMember: %v", err)
		return handler.UnauthorizedResponse()
	}

	checkResp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !checkResp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteMember: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: DeleteMember Authenticated: %v", checkResp.MemberID)

	cleaResp, err := h.members.GetClearance(context.Background(), checkResp)

	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetClearance: %v", err)
		return handler.InternalServerErrorResponse()
	}

	if !cleaResp.Access {
		h.logger.Warnf("Not allowed: %s", cleaResp.Comment)
		return handler.ForbiddenResponse()
	}

	h.logger.Infof("MembersHandler: DeleteMember Allowed: %v", checkResp.MemberID)

	memberId := &requests.DeleteMember{}

	err = memberId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind DeleteMember: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MembersHandler: parse request DeleteMember: %v", memberId)

	err = h.members.DeleteMember(context.Background(), memberId.ID)
	if err != nil {
		h.logger.Warnf("can't MembersService.DeleteMember: %v", err)
		if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else if errors.Is(err, postgres.ErrPostgresNotFoundError) {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("MembersHandler: request DeleteMember done")

	return handler.OkResponse(nil)
}

// UpdateMember updates a member's information
//
//	@Summary     Update a member's information
//	@Description Update a member's information with the provided data
//	@Tags        auth.members
//	@Accept      json
//	@Param       request body requests.UpdateMember true "Member data"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Failure     409
//	@Failure     500
//	@Router      /members/{id} [put]
//	@Security    Authorised
func (h *MembersHandler) UpdateMember(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got UpdateMember request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token UpdateMember: %v", err)
		return handler.UnauthorizedResponse()
	}

	checkResp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !checkResp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteMember: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: UpdateMember Authenticated: %v", checkResp.MemberID)

	cleaResp, err := h.members.GetClearance(context.Background(), checkResp)

	if err != nil {
		h.logger.Warnf("can't MembersHandler.GetClearance: %v", err)
		return handler.InternalServerErrorResponse()
	}

	if !cleaResp.Access {
		h.logger.Warnf("Not allowed: %s", cleaResp.Comment)
		return handler.ForbiddenResponse()
	}

	h.logger.Infof("MembersHandler: UpdateMember Allowed: %v", checkResp.MemberID)

	member := &requests.UpdateMember{}

	err = member.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind UpdateMember: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("MembersHandler: parse request UpdateMember: %v", member)

	err = h.members.UpdateMember(context.Background(), mapper.MakeRequestUpdateMember(member))
	if err != nil {
		h.logger.Warnf("can't MembersService.UpdateMember: %v", err)
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

	h.logger.Info("MembersHandler: request UpdateMember done")

	return handler.OkResponse(nil)
}

// GetClearance checks if the member is allowed to access members
//
//	@Summary     Check member`s clearance
//	@Description Check if the member is allowed to access members
//	@Tags        auth.members
//	@Success     200 {object} responses.GetClearance
//	@Failure     401
//	@Failure     500
//	@Router      /members/clearance/post/ [get]
//	@Security    Authorised
func (h *MembersHandler) GetClearance(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("MembersHandler: got GetClearance request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token GetClearance: %v", err)
		return handler.UnauthorizedResponse()
	}

	checkResp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !checkResp.Valid {
		h.logger.Warnf("can't GuardService.Check on GetClearance: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("MembersHandler: GetClearance Authenticated: %v", checkResp.MemberID)

	cleaResp, err := h.members.GetClearance(context.Background(), checkResp)

	if err != nil {
		h.logger.Warnf("can't MembersService.GetClearance: %v", err)
		return handler.InternalServerErrorResponse()
	}

	h.logger.Info("MembersHandler: GetClearance Done")

	return handler.OkResponse(cleaResp)
}
