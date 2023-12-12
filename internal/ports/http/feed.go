package http

import (
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
)

func (h *APIHandler) GetAllFeed(w http.ResponseWriter, _ *http.Request) handler.Response {
	res, err := h.feed.GetAllFeed()

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	return handler.OkResponse(res)
}

func (h *APIHandler) GetFeed(w http.ResponseWriter, r *http.Request) handler.Response {
	id, err := strconv.Atoi(r.URL.Query()["id"][0])

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}

	res, err := h.feed.GetFeed(id)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		//handler.
	}

	return handler.OkResponse(res)
}
