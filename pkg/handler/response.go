package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func SetResponseStatus(r *http.Request, status int) {
	render.Status(r, status)
}

func RenderJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.JSON(w, r, v)
}

// Response

type Response interface {
	Body() interface{}
	HTTPCode() int
}

func OkResponse(data interface{}) Response {
	return &response{body: data, code: http.StatusOK}
}

func CreatedResponse(data interface{}) Response {
	return &response{body: data, code: http.StatusCreated}
}

func InternalServerErrorResponse() Response {
	return &response{body: nil, code: http.StatusInternalServerError}
}

func BadRequestResponse() Response {
	return &response{body: nil, code: http.StatusBadRequest}
}

func NoContentResponse() Response {
	return &response{body: nil, code: http.StatusNoContent}
}

func RequestCanceledResponse() Response {
	return &response{body: nil, code: 499}
}

type response struct {
	body interface{}
	code int
}

func (r *response) Body() interface{} {
	return r.body
}

func (r *response) HTTPCode() int {
	return r.code
}

// Renderer

type Renderer interface {
	Wrap(h func(w http.ResponseWriter, r *http.Request) Response) http.HandlerFunc
}

// Basic renderer

type jsonRenderer struct{}

func NewJSONRenderer() Renderer {
	return &jsonRenderer{}
}

func (rd *jsonRenderer) Wrap(h func(w http.ResponseWriter, r *http.Request) Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(w, r)

		writeResponse(w, r, resp)
	}
}

func writeResponse(w http.ResponseWriter, r *http.Request, resp Response) {
	if body := resp.Body(); body != nil {
		SetResponseStatus(r, resp.HTTPCode())
		RenderJSON(w, r, body)
	} else {
		w.WriteHeader(resp.HTTPCode())
	}
}
