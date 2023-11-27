package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *APIHandler) GetAllFeed(w http.ResponseWriter, r *http.Request) {
	res, err := h.feed.GetAllFeed()

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (h *APIHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query()["id"][0])

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}

	res, err := h.feed.GetFeed(id)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
