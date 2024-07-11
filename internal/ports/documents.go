package http

import (
	"context"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	_ "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type DocumentsHandler struct {
	r         handler.Renderer
	documents app.DocumentsService
	logger    *log.Logger
	guard     *app.GuardService
}

func NewDocumentsHandler(r handler.Renderer, documents app.DocumentsService, logger *log.Logger, guard *app.GuardService) *DocumentsHandler {
	return &DocumentsHandler{
		r:         r,
		documents: documents,
		logger:    logger,
		guard:     guard,
	}
}

func (h *DocumentsHandler) BasePrefix() string {
	return "/documents"
}

func (h *DocumentsHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllDocuments))
	r.Get("/{id}", h.r.Wrap(h.GetDocument))
	r.Get("/search/{club_id}", h.r.Wrap(h.GetDocumentsByClubID))
	r.Post("/", h.r.Wrap(h.PostDocument))
	r.Delete("/{id}", h.r.Wrap(h.DeleteDocument))
	r.Put("/{id}", h.r.Wrap(h.UpdateDocument))

	return r
}

// GetAllDocuments retrieves all documents items
//
//	@Summary     Retrieve all documents items
//	@Description Get a list of all documents items
//	@Tags        public.documents
//	@Produce     json
//	@Success     200 {object}  responses.GetAllDocuments
//	@Failure     404
//	@Router      /documents [get]
//	@Security    public
func (h *DocumentsHandler) GetAllDocuments(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got GetAllDocuments request")

	res, err := h.documents.GetAllDocuments(context.Background())
	if err != nil {
		h.logger.Warnf("can't DocumentsService.GetAllDocuments: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request GetAllDocuments done")

	return handler.OkResponse(res)
}

// GetDocument retrieves a document item by its ID
//
//	@Summary     Retrieve document item by ID
//	@Description Get a specific document item using its ID
//	@Tags        public.documents
//	@Produce     json
//	@Param       id   path     string           true "id"
//	@Success     200  {object} responses.GetDocument
//	@Failure     400
//	@Failure     404
//	@Router      /documents/{id} [get]
//	@Security    public
func (h *DocumentsHandler) GetDocument(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got GetDocument request")

	documentId := &requests.GetDocument{}

	err := documentId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request GetDocument: %v", documentId)

	res, err := h.documents.GetDocument(context.Background(), documentId.ID)
	if err != nil {
		h.logger.Warnf("can't DocumentsService.GetDocument: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request GetDocument done")

	return handler.OkResponse(res)
}

// GetDocumentsByClubID retrieves a slice of documents by its club ID
//
//	@Summary     Retrieve documents by club ID
//	@Description Get documents corresponding to a certain club ID (0 for main page)
//	@Tags        public.documents
//	@Produce     json
//	@Param       club_id   path     string     true "club_id"
//	@Success     200  {object} responses.GetDocumentsByClubID
//	@Failure     400
//	@Failure     404
//	@Router      /documents/search/{club_id} [get]
//	@Security    public
func (h *DocumentsHandler) GetDocumentsByClubID(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got GetDocumentsByClubID request")

	clubId := &requests.GetDocumentsByClubID{}

	err := clubId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request GetDocumentsByClubID: %v", clubId)

	res, err := h.documents.GetDocumentsByClubID(context.Background(), clubId.ID)
	if err != nil {
		h.logger.Warnf("can't DocumentsService.GetDocumentsByClubID: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request GetDocumentsByClubID done")

	return handler.OkResponse(res)
}

// PostDocument creates a new document item
//
//		@Summary     Create a new document item
//		@Description Create a new document item with the provided data
//		@Tags        auth.documents
//		@Accept      json
//		@Param       request body requests.PostDocument true "Document data"
//		@Success     201
//		@Failure     400
//	 	@Failure     401
//		@Failure     404
//		@Router      /documents/ [post]
//		@Security    Authorised
func (h *DocumentsHandler) PostDocument(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got PostDocument request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token PostDocument: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on PostDocument: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("DocumentsHandler: PostDocument Authenticated: %v", resp.MemberID)

	document := &requests.PostDocument{}

	err = document.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind PostDocument: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request PostDocument: %v", document)

	err = h.documents.PostDocument(context.Background(), *mapper.MakeRequestPostDocument(*document))
	if err != nil {
		h.logger.Warnf("can't DocumentsService.PostDocument: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request PostDocument done")

	return handler.CreatedResponse(nil)
}

// DeleteDocument deletes a document item by ID
//
//	@Summary     Delete a document item by ID
//	@Description Delete a specific document item using its ID
//	@Tags        auth.documents
//	@Param       id   path     string           true "Document ID"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Router      /documents/{id} [delete]
//	@Security    Authorised
func (h *DocumentsHandler) DeleteDocument(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got DeleteDocument request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token DeleteDocument: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteDocument: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("DocumentsHandler: DeleteDocument Authenticated: %v", resp.MemberID)

	documentId := &requests.DeleteDocument{}

	err = documentId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind DeleteDocument: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request DeleteDocument: %v", documentId)

	err = h.documents.DeleteDocument(context.Background(), documentId.ID)
	if err != nil {
		h.logger.Warnf("can't DocumentsService.DeleteDocument: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request DeleteDocument done")

	return handler.OkResponse(nil)
}

// UpdateDocument updates a document item
//
//	@Summary     Update a document item
//	@Description Update an existing document item with the provided data
//	@Tags        auth.documents
//	@Accept      json
//	@Param       id   path     string           true "Document ID"
//	@Param       request body requests.UpdateDocument true "Document new data"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Router      /documents/{id} [put]
//	@Security    Authorised
func (h *DocumentsHandler) UpdateDocument(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got UpdateDocument request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token UpdateDocument: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on UpdateDocument: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("DocumentsHandler: UpdateDocument Authenticated: %v", resp.MemberID)

	document := &requests.UpdateDocument{}

	err = document.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind UpdateDocument: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request UpdateDocument: %v", document)

	err = h.documents.UpdateDocument(context.Background(), *mapper.MakeRequestUpdateDocument(*document))
	if err != nil {
		h.logger.Warnf("can't DocumentsService.UpdateDocument: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request UpdateDocument done")

	return handler.OkResponse(nil)
}
