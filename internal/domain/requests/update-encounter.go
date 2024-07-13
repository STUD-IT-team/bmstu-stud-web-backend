package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateEncounter struct {
	ID          int    `json:"id"`
	Count       string `json:"count"`
	Description string `json:"description"`
	ClubID      int    `json:"club_id"`
}

type UpdateEncounterPointer struct {
	Count       *string `json:"count"`
	Description *string `json:"description"`
	ClubID      *int    `json:"club_id"`
}

func (ev *UpdateEncounter) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := UpdateEncounterPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateEncounter.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on UpdateEncounter.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*ev = UpdateEncounter{
		Count:       *pf.Count,
		Description: *pf.Description,
		ClubID:      *pf.ClubID,
	}

	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on UpdateEncounter.Bind: %w", err)
	}

	ev.ID = id

	return ev.validate()
}

func (ev *UpdateEncounter) validate() error {
	if ev.ID == 0 {
		return fmt.Errorf("require: ID")
	}

	return nil
}

func (pf *UpdateEncounterPointer) validate() error {
	if pf.Count == nil {
		return fmt.Errorf("require: Count")
	}
	if pf.Description == nil {
		return fmt.Errorf("require: Description")
	}
	if pf.ClubID == nil {
		return fmt.Errorf("require: ClubID")
	}
	return nil
}
