package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GuardHandler struct {
	r           handler.Renderer
	guardClient grpc2.GuardClient
}

func NewGuardHandler(r handler.Renderer, conn grpc.ClientConnInterface) *GuardHandler {
	return &GuardHandler{
		r:           r,
		guardClient: grpc2.NewGuardClient(conn),
	}
}

func (h *GuardHandler) BasePrefix() string {
	return "/guard"
}

func (h *GuardHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", h.r.Wrap(h.Login))
	r.Delete("/logout", h.r.Wrap(h.Logout))
	r.Get("/check", h.r.Wrap(h.Check))

	return r
}

// Login log in by login and password
//
//	@Summary      login
//	@Description  log in by login and password
//	@Tags         login
//	@Accept       json
//	@Produce      json
//	@Param        login body string true "login"
//	@Param        password body string true "Password"
//	@Success      200  {object}  responses.LoginResponse
//	@Failure      400  {object}  handler.Response
//	@Failure      401  {object}  handler.Response
//	@Failure      500  {object}  handler.Response
//	@Router       /guard/login [post]
func (h *GuardHandler) Login(w http.ResponseWriter, r *http.Request) handler.Response {
	var req requests.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithError(err).Warnf("can't decod http request")
		return handler.BadRequestResponse()
	}

	reqGRPC := mapper.CreateGPRCRequestLogin(req)

	resGRPC, err := h.guardClient.Login(context.Background(), reqGRPC)
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return handler.BadRequestResponse()
		} else if status.Code(err) == codes.NotFound {
			return handler.UnauthorizedResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	res := mapper.CreateResponseLogin(resGRPC.AccessToken, resGRPC.Expires)

	return handler.OkResponse(res)
}

// Logout log out by session token
//
//	@Summary      logout
//	@Description  log out by session token
//	@Tags         logout
//	@Accept       json
//	@Produce      json
//	@Param        access_token body string true "session token"
//	@Success      200  {object}  responses.LoginResponse
//	@Failure      400  {object}  handler.Response
//	@Router       /guard/logout [delete]
func (h *GuardHandler) Logout(w http.ResponseWriter, r *http.Request) handler.Response {
	var req requests.LogoutRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithError(err).Warnf("can't decod http request")
		return handler.BadRequestResponse()
	}

	reqGRPC := mapper.CreateGPRCResponseLogout(req)

	_, err := h.guardClient.Logout(context.Background(), reqGRPC)
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	return handler.OkEmptyResponse()
}

// Check chekc valid access token
//
//	@Summary      check
//	@Description  check valid access token
//	@Tags         check
//	@Accept       json
//	@Produce      json
//	@Param        login body string true "login"
//	@Param        password body string true "Password"
//	@Success      200  {object}  responses.LoginResponse
//	@Failure      400  {object}  handler.Response
//	@Failure      401  {object}  handler.Response
//	@Failure      500  {object}  handler.Response
//	@Router       /guard/login [post]
func (h *GuardHandler) Check(w http.ResponseWriter, r *http.Request) handler.Response {
	var req requests.CheckRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithError(err).Warnf("can't decod http request")
		return handler.BadRequestResponse()
	}

	reqGRPC := mapper.CreateGRPCRequestCheck(req)

	resGRPC, err := h.guardClient.Check(context.Background(), reqGRPC)
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return handler.BadRequestResponse()
		} else if status.Code(err) == codes.NotFound {
			return handler.NotFoundResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	res := mapper.CreateResponseCheck(resGRPC.Valid, resGRPC.MemberID)

	return handler.OkResponse(res)
}
