package requests

import (
	"fmt"
	"net/http"
	"strconv"
)

type GetAllFeed struct {
	LastId int
	NFeed  int
}

func (f *GetAllFeed) Bind(req *http.Request) error {
	query := req.URL.Query()
	lastId, err := strconv.Atoi(query.Get("last_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi last_id on GetAllFeed.Bind: %w", err)
	}

	nFeed, err := strconv.Atoi(query.Get("n"))
	if err != nil {
		return fmt.Errorf("can't Atoi n on GetAllFeed.Bind: %w", err)
	}

	f.LastId = lastId
	f.NFeed = nFeed

	return nil
}

func (f *GetAllFeed) GetNFeedStartLastId() error {
	if f.LastId <= 0 {
		return fmt.Errorf("require: last_id")
	}

	if f.NFeed <= 0 {
		return fmt.Errorf("require: n_feed")
	}

	return nil
}
