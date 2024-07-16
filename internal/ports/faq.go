package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type FAQHandler struct {
	r      handler.Renderer
	faq    app.FAQService
	logger *log.Logger
	guard  *app.GuardService
}

func NewFAQHandler(r handler.Renderer, faq app.FAQService, logger *log.Logger, guard *app.GuardService) *FAQHandler {
	return &FAQHandler{
		r:      r,
		faq:    faq,
		logger: logger,
		guard:  guard,
	}
}

func (h *FAQHandler) BasePrefix() string {
	return "/faq"
}

func (h *FAQHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/get_all_faq/", h.r.Wrap(h.GetAllFAQ))
	r.Get("/get_by_club_id/{club_id}", h.r.Wrap(h.GetFAQByClubid))
	//r.Get("/get_by_type/{type_question}", h.r.Wrap(h.GetFAQByTypeQuestion))
	r.Post("/", h.r.Wrap(h.PostFAQ))
	r.Delete("/{id}", h.r.Wrap(h.DeleteFAQ))
	r.Put("/{id}", h.r.Wrap(h.UpdateFAQ))

	return r
}
func (h *FAQHandler) GetAllFAQ(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FaqHandler: got GetAllFAQ request")

	res, err := h.faq.GetAllFAQ(context.Background())
	if err != nil {
		h.logger.Warnf("can't FaqService.GetAllFAQ: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("FaqHandler: request GetAllFAQ done")
	return handler.OkResponse(res)
}

func (h *FAQHandler) GetFAQ(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("FAQHandler: got GetFAQ request")

	reqFAQ := &requests.GetFAQ{}

	err := reqFAQ.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("FAQHandler: parse request GetFAQ: %v", reqFAQ)

	res, err := h.faq.GetFAQ(context.Background(), reqFAQ.ID)
	if err != nil {
		h.logger.Warnf("can't FAQService.GetFAQ: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("FAQHandler: request GetFAQ done")

	return handler.OkResponse(res)
}

func (h *FAQHandler) GetFAQByClubid(w http.ResponseWriter, req *http.Request) handler.Response {
	// h.logger.Info("FAQHandler: got GetFAQByTitle request")

	// filter := &requests.GetFAQByClubid{}

	// err := filter.Bind(req)
	// if err != nil {
	// 	h.logger.Warnf("can't requests.Bind GetFAQByClubid: %v", err)
	// 	return handler.BadRequestResponse()
	// }

	// h.logger.Infof("FAQHandler: parse request GetFAQByClubid: %v", filter)

	// res, err := h.faq.GetFAQByTitle(context.Background(), *filter)
	// if err != nil {
	// 	h.logger.Warnf("can't FAQService.GetFAQByClubid: %v", err)
	// 	return handler.NotFoundResponse()
	// }

	// h.logger.Info("FAQHandler: request GetFAQByClubid done")

	return handler.OkResponse(nil)
}

func (h *FAQHandler) PostFAQ(w http.ResponseWriter, req *http.Request) handler.Response {
	// h.logger.Info("FAQHandler: got PostFAQ request")

	// accessToken, err := getAccessToken(req)
	// if err != nil {
	// 	h.logger.Warnf("can't get access token PostFAQ: %v", err)
	// 	return handler.UnauthorizedResponse()
	// }

	// resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	// if err != nil || !resp.Valid {
	// 	h.logger.Warnf("can't GuardService.Check on DeleteFAQ: %v", err)
	// 	return handler.UnauthorizedResponse()
	// }

	// h.logger.Infof("FAQHandler: PostFAQ Authenticated: %v", resp.MemberID)

	// faq := &requests.PostFAQ{}

	// err = faq.Bind(req)
	// if err != nil {
	// 	h.logger.Warnf("can't requests.Bind PostFAQ: %v", err)
	// 	return handler.BadRequestResponse()
	// }

	// h.logger.Infof("FAQHandler: parse request PostFAQ: %v", faq)

	// err = h.faq.PostFAQ(context.Background(), *mapper.MakeRequestPostFAQ(*faq))
	// if err != nil {
	// 	h.logger.Warnf("can't FAQService.PostFAQ: %v", err)
	// 	return handler.NotFoundResponse()
	// }

	// h.logger.Info("FAQHandler: request PostFAQ done")

	return handler.CreatedResponse(nil)
}

func (h *FAQHandler) DeleteFAQ(w http.ResponseWriter, req *http.Request) handler.Response {
	// h.logger.Info("FAQHandler: got DeleteFAQ request")

	// accessToken, err := getAccessToken(req)
	// if err != nil {
	// 	h.logger.Warnf("can't get access token DeleteFAQ: %v", err)
	// 	return handler.UnauthorizedResponse()
	// }

	// resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	// if err != nil || !resp.Valid {
	// 	h.logger.Warnf("can't GuardService.Check on DeleteFAQ: %v", err)
	// 	return handler.UnauthorizedResponse()
	// }

	// h.logger.Infof("FAQHandler: DeleteFAQ Authenticated: %v", resp.MemberID)

	// faqId := &requests.DeleteFAQ{}

	// err = faqId.Bind(req)
	// if err != nil {
	// 	h.logger.Warnf("can't requests.Bind DeleteFAQ: %v", err)
	// 	return handler.BadRequestResponse()
	// }

	// h.logger.Infof("FAQHandler: parse request DeleteFAQ: %v", faqId)

	// err = h.faq.DeleteFAQ(context.Background(), faqId.ID)
	// if err != nil {
	// 	h.logger.Warnf("can't FAQService.DeleteFAQ: %v", err)
	// 	return handler.NotFoundResponse()
	// }

	// h.logger.Info("FAQHandler: request DeleteFAQ done")

	return handler.OkResponse(nil)
}

func (h *FAQHandler) UpdateFAQ(w http.ResponseWriter, req *http.Request) handler.Response {
	// h.logger.Info("FAQHandler: got UpdateFAQ request")

	// accessToken, err := getAccessToken(req)
	// if err != nil {
	// 	h.logger.Warnf("can't get access token UpdateFAQ: %v", err)
	// 	return handler.UnauthorizedResponse()
	// }

	// resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	// if err != nil || !resp.Valid {
	// 	h.logger.Warnf("can't GuardService.Check on DeleteFAQ: %v", err)
	// 	return handler.UnauthorizedResponse()
	// }

	// h.logger.Infof("FAQHandler: UpdateFAQ Authenticated: %v", resp.MemberID)

	// faq := &requests.UpdateFAQ{}

	// err = faq.Bind(req)
	// if err != nil {
	// 	h.logger.Warnf("can't requests.Bind UpdateFAQ: %v", err)
	// 	return handler.BadRequestResponse()
	// }

	// h.logger.Infof("FAQHandler: parse request UpdateFAQ: %v", faq)

	// err = h.faq.UpdateFAQ(context.Background(), *mapper.MakeRequestUpdateFAQ(*&faq))
	// if err != nil {
	// 	h.logger.Warnf("can't FAQService.UpdateFAQ: %v", err)
	// 	return handler.NotFoundResponse()
	// }

	// h.logger.Info("FAQHandler: request UpdateFAQ done")

	return handler.OkResponse(nil)
}
