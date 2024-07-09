package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostEvent struct {
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Propmt      string    `db:"propmt"`
	MediaID     int       `db:"media_id"`
	Date        time.Time `db:"date"`
	Approved    bool      `db:"approved"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedBy   int       `db:"created_by"`
	RegUrl      string    `db:"reg_url"`
	RegOpenDate time.Time `db:"reg_open_date"`
	FeedbackUrl string    `db:"eventback_url"`
}

type PostEventPointer struct {
	Title       *string    `db:"title"`
	Description *string    `db:"description"`
	Propmt      *string    `db:"propmt"`
	MediaID     *int       `db:"media_id"`
	Date        *time.Time `db:"date"`
	Approved    *bool      `db:"approved"`
	CreatedAt   *time.Time `db:"created_at"`
	CreatedBy   *int       `db:"created_by"`
	RegUrl      *string    `db:"reg_url"`
	RegOpenDate *time.Time `db:"reg_open_date"`
	FeedbackUrl *string    `db:"eventback_url"`
}

func (f *PostEvent) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := PostEventPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostEvent.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostEvent.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*f = PostEvent{
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

	return f.validate()
}

func (f *PostEvent) validate() error {
	return nil
}

func (pf *PostEventPointer) validate() error {
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