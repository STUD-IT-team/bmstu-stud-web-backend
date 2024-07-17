package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type DeleteClubPhoto struct {
	PhotoID int `json:"photo_id"`
	ClubID  int `json:"club_id"`
}

type DeleteClubPhotoPointer struct {
	ClubID  int  `json:"club_id"`
	PhotoID *int `json:"photo_id"`
}

func (c *DeleteClubPhoto) Bind(req *http.Request) error {
	clubID, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi on club_id in request: %w", err)
	}

	pc := DeleteClubPhotoPointer{}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&pc)
	if err != nil {
		return fmt.Errorf("can't json decoder on DeleteClubPhoto: %v", err)
	}
	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on DeleteClubPhoto")
	}

	err = pc.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	c.ClubID = clubID
	c.PhotoID = *pc.PhotoID
	return c.validate()
}

func (c *DeleteClubPhoto) validate() error {
	if c.ClubID == 0 {
		return fmt.Errorf("require: id")
	}
	return nil
}

func (pc *DeleteClubPhotoPointer) validate() error {
	if pc.PhotoID == nil {
		return fmt.Errorf("require: photo_id")
	}
	return nil
}
