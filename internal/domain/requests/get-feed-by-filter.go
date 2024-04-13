package requests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/samber/mo"
)

type GetFeedByFilter struct {
	Offset int
	Limit  int
}

func (f *GetFeedByFilter) Bind(req *http.Request) error {
	query := req.URL.Query()

	var limit, offset int
	var err error

	stringOffset := query.Get("offset")
	if optionOffset := mo.TupleToOption(stringOffset, stringOffset != ""); optionOffset.IsPresent() {
		offset, err = strconv.Atoi(optionOffset.MustGet())
		if err != nil {
			return fmt.Errorf("can't Atoi offset on GetAllFeed.Bind: %w", err)
		}
	}

	stringLimit := query.Get("limit")
	if optionLimit := mo.TupleToOption(stringLimit, stringLimit != ""); optionLimit.IsPresent() {
		limit, err = strconv.Atoi(optionLimit.MustGet())
		if err != nil {
			return fmt.Errorf("can't Atoi limit on GetAllFeed.Bind: %w", err)
		}
	}

	f.Offset = offset
	f.Limit = limit

	return f.GetLimitOffset()
}

func (f *GetFeedByFilter) GetLimitOffset() error {
	if f.Offset < 0 {
		return fmt.Errorf("require: offset <= 0")
	}

	if f.Limit < 0 {
		return fmt.Errorf("require: limit <= 0")
	}

	return nil
}
