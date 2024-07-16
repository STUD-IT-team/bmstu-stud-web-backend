package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"
)

type PostClubPhoto struct {
	ClubID int         `json:"club_id"`
	Photos []ClubPhoto `json:"photos"`
}

type PostClubPhotoPointer struct {
	ClubID int         `json:"club_id"`
	Photos []ClubPhoto `json:"photos"`
}

type ClubPhoto struct {
	MediaID   int `json:"media_id"`
	RefNumber int `json:"ref_number"`
}

func (c *PostClubPhoto) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "club_id"))
	if err != nil {
		return fmt.Errorf("can't Atoi on club_id in request: %w", err)
	}
	pc := PostClubPhotoPointer{}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&pc)

	if err != nil {
		return fmt.Errorf("can't json decoder on PostClub: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("postClub Bind: extraneous data after JSON object")
	}

	err = pc.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}
	*c = PostClubPhoto{
		ClubID: id,
		Photos: pc.Photos,
	}
	return c.validate()

}

func (c *PostClubPhoto) validate() error {
	return nil
}

func (pc *PostClubPhotoPointer) validate() error {
	if pc.Photos == nil {
		return fmt.Errorf("require: photos")
	}
	return nil
}
