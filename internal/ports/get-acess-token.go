package http

import (
	"fmt"
	"net/http"
)

func getAccessToken(req *http.Request) (string, error) {
	c, err := req.Cookie("AccessToken")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", fmt.Errorf("cookie is not set PostFeed: %v", err)
		}
		return "", fmt.Errorf("cookie error PostFeed: %v", err)
	}
	return c.Value, nil
}
