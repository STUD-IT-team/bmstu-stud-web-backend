package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostMedia struct {
	Name string
	Data []byte
}

type PostMediaPointer struct {
	Name *string
	Data []byte
}

const maxMemory = 10 << 24

func (m *PostMedia) Bind(req *http.Request) error {

	if err := req.ParseMultipartForm(maxMemory); err != nil {
		return fmt.Errorf("can't parse multipart form on PostMedia.Bind: %v", err)
	}

	pm := PostMediaPointer{}

	// name
	if name := req.FormValue("name"); name != "" {
		pm.Name = &name
	}

	// data (file)
	file, _, err := req.FormFile("data")
	if err != nil {
		return fmt.Errorf("%v: require: Data", domain.ErrIncorrectRequest)
	}
	defer file.Close()

	pm.Data, err = io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("can't read file data on PostMedia.Bind: %v", err)
	}

	if err := pm.validate(); err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*m = PostMedia{
		Name: *pm.Name,
		Data: pm.Data,
	}

	return m.validate()
}

func (m *PostMedia) validate() error {
	return nil
}

func (pm *PostMediaPointer) validate() error {
	if pm.Name == nil {
		return fmt.Errorf("require: Name")
	}
	if pm.Data == nil {
		return fmt.Errorf("require: Data")
	}
	return nil
}
