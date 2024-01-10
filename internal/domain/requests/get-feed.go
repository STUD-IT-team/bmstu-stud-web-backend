package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetFeedBind(req *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(req, "id"))
}
