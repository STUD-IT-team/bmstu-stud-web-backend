package postgres

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

const getMediaFile = "SELECT name, key FROM mediafile WHERE id = $1"

func (p *Postgres) GetMediaFile(id int) (*domain.MediaFile, error) {
	f := domain.MediaFile{}
	err := p.db.QueryRow(getMediaFile, id).Scan(&f.Name, &f.Key)
	if err == nil {
		return &f, nil
	}
	f.ID = id
	return nil, err
}

const getMediaFiles = "SELECT id, name, key FROM mediafile WHERE id = ANY($1)"

func (p *Postgres) GetMediaFiles(ids []int) (map[int]domain.MediaFile, error) {
	m := make(map[int]domain.MediaFile)
	rows, err := p.db.Query(getMediaFiles, ids)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		media := domain.MediaFile{}
		err := rows.Scan(&media.ID, &media.Name, &media.Key)
		if err != nil {
			return nil, err
		}
		m[media.ID] = media
	}
	return m, nil
}

const addMediaFile = "INSERT INTO mediafile (name, key) VALUES ($1, $2) RETURNING id"

func (p *Postgres) AddMediaFile(name, key string) (int, error) {
	var id int
	err := p.db.QueryRow(addMediaFile, name, key).Scan(&id)
	return id, err
}
