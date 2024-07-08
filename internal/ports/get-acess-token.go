package http

import (
	"fmt"
	"net/http"
	"strconv"
)

func getAccessToken(req *http.Request) (int64, error) {
	c, err := req.Cookie("AccessToken")
	if err != nil {
		if err == http.ErrNoCookie {
			return 0, fmt.Errorf("cookie is not set PostFeed: %v", err)
		}
		return 0, fmt.Errorf("cookie error PostFeed: %v", err)
	}
	id, err := strconv.ParseInt(c.Value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("can't ParseInt on cookie value: %w", err)
	}
	return id, nil
}
