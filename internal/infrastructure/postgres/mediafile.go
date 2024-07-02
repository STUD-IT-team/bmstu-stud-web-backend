package postgres

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

const getMediaFile = "SELECT name, image FROM mediafile WHERE id = $1"

func (p *Postgres) GetMediaFile(id int) (*domain.MediaFile, error) {
	f := domain.MediaFile{}
	err := p.db.QueryRow(getMediaFile, id).Scan(&f.Name, &f.Image)
	if err == nil {
		return &f, nil
	}
	return nil, err
}
