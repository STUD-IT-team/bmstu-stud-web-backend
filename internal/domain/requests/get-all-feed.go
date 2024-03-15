package requests

import (
	"fmt"
	"net/http"
	"strconv"
)

type GetAllFeed struct {
	Offset int
	Limit  int
}

func (f *GetAllFeed) Bind(req *http.Request) error {
	query := req.URL.Query()

	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		return fmt.Errorf("can't Atoi offset on GetAllFeed.Bind: %w", err)
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		return fmt.Errorf("can't Atoi limit on GetAllFeed.Bind: %w", err)
	}

	f.Offset = offset
	f.Limit = limit

	return nil
}

func (f *GetAllFeed) GetLimitOffset() error {
	if f.Offset <= 0 {
		return fmt.Errorf("require: offset")
	}

	if f.Limit <= 0 {
		return fmt.Errorf("require: limit")
	}

	return nil
}
