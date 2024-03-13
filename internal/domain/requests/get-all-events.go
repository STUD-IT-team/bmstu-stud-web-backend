package requests

import "time"

type GetAllEvents struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (e *GetAllEvents) validate() error  {
	if e.EndDate
}