package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostEncounter struct {
	Count       string `json:"count"`
	Description string `json:"description"`
	ClubID      int    `json:"club_id"`
}

type PostEncounterPointer struct {
	Count       *string `json:"count"`
	Description *string `json:"description"`
	ClubID      *int    `json:"club_id"`
}

func (f *PostEncounter) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := PostEncounterPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostEncounter.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostEncounter.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*f = PostEncounter{
		Count:       *pf.Count,
		Description: *pf.Description,
		ClubID:      *pf.ClubID,
	}

	return f.validate()
}

func (f *PostEncounter) validate() error {
	return nil
}

func (pf *PostEncounterPointer) validate() error {
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
