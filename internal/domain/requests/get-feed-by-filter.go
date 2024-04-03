package requests

import (
	"fmt"
	"net/http"
	"strconv"
)

type GetFeedByFilter struct {
	Offset int
	Limit  int
}

func (f *GetFeedByFilter) Bind(req *http.Request) error {
	var limit, offset int
	var err error
	query := req.URL.Query()

	if query.Has("offset") {
		offset, err = strconv.Atoi(query.Get("offset"))
		if err != nil {
			return fmt.Errorf("can't Atoi offset on GetAllFeed.Bind: %w", err)
		}
		if offset < 0 {
			return fmt.Errorf("require: offset < 0")
		}
	}

	if query.Has("limit") {
		limit, err = strconv.Atoi(query.Get("limit"))
		if err != nil {
			return fmt.Errorf("can't Atoi limit on GetAllFeed.Bind: %w", err)
		}
		if limit < 0 {
			return fmt.Errorf("require: limit < 0")
		}
	}

	f.Offset = offset
	f.Limit = limit

	return nil
}

func (f *GetFeedByFilter) GetLimitOffset() error {
	if f.Offset <= 0 {
		return fmt.Errorf("require: offset <= 0")
	}

	if f.Limit <= 0 {
		return fmt.Errorf("require: limit <= 0")
	}

	return nil
}
