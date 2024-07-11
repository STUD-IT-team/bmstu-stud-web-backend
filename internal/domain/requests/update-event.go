package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateEvent struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Propmt      string    `json:"propmt"`
	MediaID     int       `json:"media_id"`
	Date        time.Time `json:"date"`
	Approved    bool      `json:"approved"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	RegUrl      string    `json:"reg_url"`
	RegOpenDate time.Time `json:"reg_open_date"`
	FeedbackUrl string    `json:"eventback_url"`
}

type UpdateEventPointer struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Propmt      *string    `json:"propmt"`
	MediaID     *int       `json:"media_id"`
	Date        *time.Time `json:"date"`
	Approved    *bool      `json:"approved"`
	CreatedAt   *time.Time `json:"created_at"`
	CreatedBy   *int       `json:"created_by"`
	RegUrl      *string    `json:"reg_url"`
	RegOpenDate *time.Time `json:"reg_open_date"`
	FeedbackUrl *string    `json:"eventback_url"`
}

func (ev *UpdateEvent) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := UpdateEventPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateEvent.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on UpdateEvent.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*ev = UpdateEvent{
		Title:       *pf.Title,
		Description: *pf.Description,
		Propmt:      *pf.Propmt,
		MediaID:     *pf.MediaID,
		Date:        *pf.Date,
		Approved:    *pf.Approved,
		CreatedAt:   *pf.CreatedAt,
		CreatedBy:   *pf.CreatedBy,
		RegUrl:      *pf.RegUrl,
		RegOpenDate: *pf.RegOpenDate,
		FeedbackUrl: *pf.FeedbackUrl,
	}

	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on UpdateFeed.Bind: %w", err)
	}

	ev.ID = id

	return ev.validate()
}

func (ev *UpdateEvent) validate() error {
	return nil
}

func (pf *UpdateEventPointer) validate() error {
	if pf.Title == nil {
		return fmt.Errorf("require: Title")
	}
	if pf.Description == nil {
		return fmt.Errorf("require: Description")
	}
	if pf.Propmt == nil {
		return fmt.Errorf("require: Propmt")
	}
	if pf.MediaID == nil {
		return fmt.Errorf("require: MediaID")
	}
	if pf.Date == nil {
		return fmt.Errorf("require: Date")
	}
	if pf.Approved == nil {
		return fmt.Errorf("require: Approved")
	}
	if pf.CreatedAt == nil {
		return fmt.Errorf("require: CreatedAt")
	}
	if pf.CreatedBy == nil {
		return fmt.Errorf("require: CreatedBy")
	}
	if pf.RegUrl == nil {
		return fmt.Errorf("require: RegUrl")
	}
	if pf.RegOpenDate == nil {
		return fmt.Errorf("require: RegOpenDate")
	}
	if pf.FeedbackUrl == nil {
		return fmt.Errorf("require: FeedbackUrl")
	}
	return nil
}
