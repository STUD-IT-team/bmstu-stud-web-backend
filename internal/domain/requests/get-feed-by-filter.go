package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/samber/mo"
)

type GetFeedByFilter struct {
	Offset mo.Option[int]
	Limit  mo.Option[int]
}

func (f *GetFeedByFilter) Bind(req *http.Request) error {
	query := req.URL.Query()

	if query.Has("offset") {
		offset, err := strconv.Atoi(query.Get("offset"))
		if err != nil {
			return fmt.Errorf("can't Atoi offset on GetAllFeed.Bind: %w", err)
		}

		f.Offset = mo.Some(offset)
	}

	if query.Has("limit") {
		limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			return fmt.Errorf("can't Atoi limit on GetAllFeed.Bind: %w", err)
		}

		f.Limit = mo.Some(limit)
	}

	return f.ParseQueryParam()
}

func (f *GetFeedByFilter) ParseQueryParam() error {
	if f.Limit.IsPresent() && f.Offset.IsPresent() { // get feeds with filters: limit, offset
		return nil
	}

	if f.Limit.IsAbsent() && f.Offset.IsAbsent() { // get all feeds
		return nil
	}

	return fmt.Errorf("request doesnt exist")
}
