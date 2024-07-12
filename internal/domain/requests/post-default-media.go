package requests

import (
	"encoding/json"
	"net/http"
)

type PostDefaultMedia struct {
	PostMedia
}

type PostMediaDefaultPointer struct {
	PostMediaPointer
}

func (m *PostDefaultMedia) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pm := PostMediaDefaultPointer{}
	err := decoder.Decode(&pm)
	if err != nil {
		return err
	}
	if decoder.More() {
		return err
	}
	err = pm.validate()
	if err != nil {
		return err
	}
	*m = PostDefaultMedia{
		PostMedia: PostMedia{
			Name: *pm.Name,
			Data: pm.Data,
		},
	}
	return m.validate()
}

func (pf *PostMediaDefaultPointer) validate() error {
	return pf.PostMediaPointer.validate()
}

func (m *PostDefaultMedia) validate() error {
	return m.PostMedia.validate()
}
