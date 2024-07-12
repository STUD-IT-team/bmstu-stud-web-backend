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

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type DocumentsHandler struct {
	r          handler.Renderer
	documents  app.DocumentsService
	categories app.DocumentCategoriesService
	logger     *log.Logger
	guard      *app.GuardService
}

func NewDocumentsHandler(r handler.Renderer, documents app.DocumentsService, categories app.DocumentCategoriesService, logger *log.Logger, guard *app.GuardService) *DocumentsHandler {
	return &DocumentsHandler{
		r:          r,
		documents:  documents,
		categories: categories,
		logger:     logger,
		guard:      guard,
	}
}

func (h *DocumentsHandler) BasePrefix() string {
	return "/documents"
}

func (h *DocumentsHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllDocuments))
	r.Get("/{id}", h.r.Wrap(h.GetDocument))
	r.Get("/search_category/{category_id}", h.r.Wrap(h.GetDocumentsByCategory))
	r.Get("/search_club/{club_id}", h.r.Wrap(h.GetDocumentsByClubID))
	r.Get("/categories/", h.r.Wrap(h.GetAllCategories))
	r.Get("/categories/{id}", h.r.Wrap(h.GetCategory))
	r.Post("/", h.r.Wrap(h.PostDocument))
	r.Delete("/{id}", h.r.Wrap(h.DeleteDocument))
	r.Put("/{id}", h.r.Wrap(h.UpdateDocument))
	r.Post("/categories/", h.r.Wrap(h.PostCategory))
	r.Delete("/categories/{id}", h.r.Wrap(h.DeleteCategory))
	r.Put("/categories/{id}", h.r.Wrap(h.UpdateCategory))

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
//	@Router      /documents/ [get]
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

// GetDocumentsByCategory retrieves a slice of documents by its category
//
//	@Summary     Retrieve documents by category
//	@Description Get documents corresponding to a certain category
//	@Tags        public.documents
//	@Produce     json
//	@Param       category_id   path     string     true "category_id"
//	@Success     200  {object} responses.GetDocumentsByCategory
//	@Failure     400
//	@Failure     404
//	@Router      /documents/search_category/{category_id} [get]
//	@Security    public
func (h *DocumentsHandler) GetDocumentsByCategory(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got GetDocumentsByCategory request")

	categoryId := &requests.GetDocumentsByCategory{}

	err := categoryId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request GetDocumentsByCategory: %v", categoryId)

	res, err := h.documents.GetDocumentsByCategory(context.Background(), categoryId.ID)
	if err != nil {
		h.logger.Warnf("can't DocumentsService.GetDocumentsByCategory: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request GetDocumentsByCategory done")

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
//	@Router      /documents/search_club/{club_id} [get]
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

// GetAllCategories retrieves all available categories of documents
//
//	@Summary     Retrieve all document categories
//	@Description Get a list of all categories to which a document may be attributed to
//	@Tags        public.documents
//	@Produce     json
//	@Success     200 {object}  responses.GetAllDocumentCategories
//	@Failure     404
//	@Router      /documents/categories/ [get]
//	@Security    public
func (h *DocumentsHandler) GetAllCategories(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got GetAllCategories request")

	res, err := h.categories.GetAllDocumentCategories(context.Background())
	if err != nil {
		h.logger.Warnf("can't DocumentsService.GetAllCategories: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request GetAllCategories done")

	return handler.OkResponse(res)
}

// GetCategory retrieves a document category by its ID
//
//	@Summary     Retrieve document category by ID
//	@Description Get a specific document category using its ID
//	@Tags        public.documents
//	@Produce     json
//	@Param       id   path     string           true "id"
//	@Success     200  {object} responses.GetDocumentCategory
//	@Failure     400
//	@Failure     404
//	@Router      /documents/categories/{id} [get]
//	@Security    public
func (h *DocumentsHandler) GetCategory(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got GetCategory request")

	catId := &requests.GetDocumentCategory{}

	err := catId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request GetCategory: %v", catId)

	res, err := h.categories.GetDocumentCategory(context.Background(), catId.ID)
	if err != nil {
		h.logger.Warnf("can't DocumentsService.GetCategory: %v", err)
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request GetCategory done")

	return handler.OkResponse(res)
}

// PostDocument creates a new document item
//
//		@Summary     Create a new document item
//		@Description Create a new document item with the provided data and returns the key for it
//		@Tags        auth.documents
//		@Accept      json
//		@Produce     json
//		@Param       request body requests.PostDocument true "Document data"
//		@Success     201 {object} responses.PostDocument
//		@Failure     400
//	 	@Failure     401
//		@Failure     409
//		@Failure     500
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

	doc := &requests.PostDocument{}

	err = doc.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind PostDocument: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request PostDocument")

	res, err := h.documents.PostDocument(context.Background(), doc)
	if err != nil {
		h.logger.Warnf("can't DocumentsService.PostDocument: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("DocumentsHandler: request PostDocument done")

	return handler.CreatedResponse(res)
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
//	@Description Update an existing document item with the provided data and returns the new key for it
//	@Tags        auth.documents
//	@Accept      json
//	@Produce     json
//	@Param       id   path     string           true "Document ID"
//	@Param       request body requests.PostDocument true "Document new data"
//	@Success     200 {object} responses.UpdateDocument
//	@Failure     400
//	@Failure     401
//	@Failure     409
//	@Failure     500
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

	doc := &requests.UpdateDocument{}

	err = doc.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind UpdateDocument: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request UpdateDocument")

	res, err := h.documents.UpdateDocument(context.Background(), doc)
	if err != nil {
		h.logger.Warnf("can't DocumentsService.UpdateDocument: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("DocumentsHandler: request UpdateDocument done")

	return handler.OkResponse(res)
}

// PostCategory creates a new document category
//
// @Summary     Create a new document category
// @Description Create a new category for documents with given name
// @Tags        auth.documents
// @Accept      json
// @Param       request body requests.PostDocumentCategory true "DocumentCategory data"
// @Success     201
// @Failure     400
// @Failure     401
// @Failure     409
// @Failure     500
// @Router      /documents/categories/ [post]
// @Security    Authorised
func (h *DocumentsHandler) PostCategory(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got PostCategory request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token PostCategory: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on PostCategory: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("DocumentsHandler: PostCategory Authenticated: %v", resp.MemberID)

	cat := &requests.PostDocumentCategory{}

	err = cat.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind PostDocumentCategory: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request PostCategory")

	err = h.categories.PostDocumentCategory(context.Background(), mapper.MakeRequestPostDocumentCategory(cat))
	if err != nil {
		h.logger.Warnf("can't DocumentCategoriesService.PostDocumentCategory: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("DocumentsHandler: request PostCategory done")

	return handler.CreatedResponse(nil)
}

// DeleteCategory deletes a document category by ID
//
//	@Summary     Delete a document category by ID
//	@Description Delete a specific document category using its ID
//	@Tags        auth.documents
//	@Param       id   path     string           true "DocumentCategory ID"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     404
//	@Router      /documents/categories/{id} [delete]
//	@Security    Authorised
func (h *DocumentsHandler) DeleteCategory(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got DeleteCategory request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token DeleteDocumentCategory: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on DeleteDocumentCategory: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("DocumentsHandler: DeleteCategory Authenticated: %v", resp.MemberID)

	catId := &requests.DeleteDocumentCategory{}

	err = catId.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind DeleteDocumentCategory: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request DeleteCategory: %v", catId)

	err = h.categories.DeleteDocumentCategory(context.Background(), catId.ID)
	if err != nil {
		h.logger.Warnf("can't DocumentCategoriesService.DeleteDocumentCategory: %v", err)
		if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		}
		return handler.NotFoundResponse()
	}

	h.logger.Info("DocumentsHandler: request DeleteCategory done")

	return handler.OkResponse(nil)
}

// UpdateCategory updates a document category
//
//	@Summary     Update a document category
//	@Description Update an existing document category with the provided data
//	@Tags        auth.documents
//	@Accept      json
//	@Param       id   path     string           true "DocumentCategory ID"
//	@Param       request body requests.PostDocumentCategory true "DocumentCategory new data"
//	@Success     200
//	@Failure     400
//	@Failure     401
//	@Failure     409
//	@Failure     500
//	@Router      /documents/categories/{id} [put]
//	@Security    Authorised
func (h *DocumentsHandler) UpdateCategory(w http.ResponseWriter, req *http.Request) handler.Response {
	h.logger.Info("DocumentsHandler: got UpdateCategory request")

	accessToken, err := getAccessToken(req)
	if err != nil {
		h.logger.Warnf("can't get access token UpdateDocumentCategory: %v", err)
		return handler.UnauthorizedResponse()
	}

	resp, err := h.guard.Check(context.Background(), &requests.CheckRequest{AccessToken: accessToken})
	if err != nil || !resp.Valid {
		h.logger.Warnf("can't GuardService.Check on UpdateDocumentCategory: %v", err)
		return handler.UnauthorizedResponse()
	}

	h.logger.Infof("DocumentsHandler: UpdateCategory Authenticated: %v", resp.MemberID)

	cat := &requests.UpdateDocumentCategory{}

	err = cat.Bind(req)
	if err != nil {
		h.logger.Warnf("can't requests.Bind UpdateDocumentCategory: %v", err)
		return handler.BadRequestResponse()
	}

	h.logger.Infof("DocumentsHandler: parse request UpdateCategory")

	err = h.categories.UpdateDocumentCategory(context.Background(), mapper.MakeRequestUpdateDocumentCategory(cat))
	if err != nil {
		h.logger.Warnf("can't DocumentCategoriesService.UpdateDocumentCategory: %v", err)
		if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
			return handler.ConflictResponse()
		} else if errors.Is(err, postgres.ErrPostgresForeignKeyViolation) {
			return handler.BadRequestResponse()
		} else {
			return handler.InternalServerErrorResponse()
		}
	}

	h.logger.Info("DocumentsHandler: request UpdateCategory done")

	return handler.OkResponse(nil)
}
