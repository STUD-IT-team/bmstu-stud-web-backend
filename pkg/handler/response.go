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

const ErrorHeaderKey = "Error"

type Response interface {
	SetKVHeader(k, v string)
	Head() map[string]string
	Body() interface{}
	HTTPCode() int
}

func OkResponse(data interface{}) Response {
	return &response{head: map[string]string{}, body: data, code: http.StatusOK}
}

func CreatedResponse(data interface{}) Response {
	return &response{head: map[string]string{}, body: data, code: http.StatusCreated}
}

func NotFoundResponse() Response {
	return &response{head: map[string]string{}, body: nil, code: http.StatusNotFound}
}

func InternalServerErrorResponse() Response {
	return &response{head: map[string]string{}, body: nil, code: http.StatusInternalServerError}
}

func BadRequestResponse() Response {
	return &response{head: map[string]string{}, body: nil, code: http.StatusBadRequest}
}

func UnauthorizedResponse() Response {
	return &response{head: map[string]string{}, body: nil, code: http.StatusUnauthorized}
}

func NoContentResponse() Response {
	return &response{head: map[string]string{}, body: nil, code: http.StatusNoContent}
}

func RequestCanceledResponse() Response {
	return &response{head: map[string]string{}, body: nil, code: 499}
}

type response struct {
	head map[string]string
	body interface{}
	code int
}

func (r *response) Body() interface{} {
	return r.body
}

func (r *response) Head() map[string]string {
	return r.head
}

func (r *response) HTTPCode() int {
	return r.code
}

func (r *response) SetKVHeader(k, v string) {
	if r.head == nil {
		r.head = map[string]string{}
	}
	r.head[k] = v
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
	for k, v := range resp.Head() {
		w.Header().Set(k, v)
	}
	if body := resp.Body(); body != nil {
		SetResponseStatus(r, resp.HTTPCode())
		RenderJSON(w, r, body)
	} else {
		w.WriteHeader(resp.HTTPCode())
	}
}
