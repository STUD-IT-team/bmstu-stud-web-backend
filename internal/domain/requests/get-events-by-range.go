package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type GetEventsByRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type GetEventsByRangePointer struct {
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

func (ev *GetEventsByRange) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pev := GetEventsByRangePointer{}

	err := decoder.Decode(&pev)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostFeed.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostFeed.Bind")
	}

	err = pev.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*ev = GetEventsByRange{
		From: *pev.From,
		To:   *pev.To,
	}

	return ev.validate()
}

func (ev *GetEventsByRange) validate() error {
	if ev.From.After(ev.To) {
		return fmt.Errorf("invalid time period: 'from' must be before 'to'")
	}

	return nil
}

func (pev *GetEventsByRangePointer) validate() error {
	if pev.From == nil {
		return fmt.Errorf("require: From")
	}
	if pev.To == nil {
		return fmt.Errorf("require: To")
	}

	return nil
}
