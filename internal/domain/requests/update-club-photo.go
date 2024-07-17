package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type UpdateClubPhoto struct {
	PostClubPhoto
}

type UpdateClubPhotoPointer struct {
	PostClubPhotoPointer
}

func (c *UpdateClubPhoto) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi on club_id in request: %w", err)
	}
	pc := PostClubPhotoPointer{}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&pc)

	if err != nil {
		return fmt.Errorf("can't json decoder on updateClubPhoto: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("updateClubPhoto Bind: extraneous data after JSON object")
	}

	err = pc.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}
	c.ClubID = id
	c.Photos = pc.Photos

	return c.validate()
}
