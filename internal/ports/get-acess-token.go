package http

import (
	"fmt"
	"net/http"
	"strconv"
)

func getAccessToken(req *http.Request) (int64, error) {
	// c, err := req.Cookie("AccessToken")
	c, ok := req.Header["Authorization"]
	if !ok {
		return 0, fmt.Errorf("no header Authorization")
	}
	if len(c) != 1 {
		return 0, fmt.Errorf("more than one header Authorization")
	}
	val := c[0]
	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("can't ParseInt on cookie value: %w", err)
	}
	return id, nil
}
