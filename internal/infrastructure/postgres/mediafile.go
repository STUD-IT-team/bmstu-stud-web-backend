package postgres

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/jackc/pgx"
)

const getMediaFile = "SELECT name, key FROM mediafile WHERE id = $1 AND id > 0"

func (p *Postgres) GetMediaFile(id int) (*domain.MediaFile, error) {
	f := domain.MediaFile{}
	err := p.db.QueryRow(getMediaFile, id).Scan(&f.Name, &f.Key)
	if err != nil {
		return nil, err
	}
	f.ID = id
	return &f, nil
}

const getMediaFiles = "SELECT id, name, key FROM mediafile WHERE id = ANY($1) AND id > 0"

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
	if err != nil {
		return 0, wrapPostgresError(err.(pgx.PgError).Code, err)
	}
	return id, nil
}

const deleteMediaFile = "DELETE FROM mediafile WHERE id = $1 AND id > 0"

func (p *Postgres) DeleteMediaFile(id int) error {
	_, err := p.db.Exec(deleteMediaFile, id)
	if err != nil {
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}
	return nil
}

const getUnusedMedia = `
SELECT
    id,
	name,
	key
FROM mediafile
WHERE id NOT IN (
    SELECT media_id FROM club_photo
    UNION ALL
	SELECT media_id FROM member
	UNION ALL
	SELECT media_id FROM feed
	UNION ALL
	SELECT media_id FROM event
	UNION ALL
	SELECT logo as media_id FROM club
) AND id > 0
`

func (p *Postgres) GetUnusedMedia(ctx context.Context) ([]domain.MediaFile, error) {
	res := make([]domain.MediaFile, 0)
	rows, err := p.db.Query(getUnusedMedia)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		med := domain.MediaFile{}
		err := rows.Scan(&med.ID, &med.Name, &med.Key)
		if err != nil {
			return nil, err
		}
		res = append(res, med)
	}
	return res, nil
}

const deleteMediaFiles = "DELETE FROM mediafile WHERE key = ANY($1) AND id > 0"

func (p *Postgres) DeleteMediaFiles(ctx context.Context, keys []string) error {
	_, err := p.db.Exec(deleteMediaFiles, keys)
	if err != nil {
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}
	return nil
}

const getAllMediaKeys = "SELECT key FROM mediafile WHERE id > 0"

func (p *Postgres) GetAllMediaKeys(ctx context.Context) ([]string, error) {
	rows, err := p.db.Query(getAllMediaKeys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []string
	for rows.Next() {
		var key string
		err := rows.Scan(&key)
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}

const getDefautlMedia = "SELECT id, media_id FROM default_media WHERE id = $1"

func (p *Postgres) GetDefautlMedia(ctx context.Context, id int) (*domain.DefaultMedia, error) {
	var d domain.DefaultMedia
	err := p.db.QueryRow(getDefautlMedia, id).Scan(&d.ID, &d.MediaID)
	if err == nil {
		return &d, nil
	}
	return nil, err
}

const getAllDefaultMedia = "SELECT id, media_id FROM default_media"

func (p *Postgres) GetAllDefaultMedia(ctx context.Context) ([]domain.DefaultMedia, error) {
	defaultMedia := make([]domain.DefaultMedia, 0)
	rows, err := p.db.Query(getAllDefaultMedia)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		d := domain.DefaultMedia{}
		err := rows.Scan(&d.ID, &d.MediaID)
		if err != nil {
			return nil, err
		}
		defaultMedia = append(defaultMedia, d)
	}

	return defaultMedia, nil
}

const addDefaultMedia = "INSERT INTO default_media (media_id) VALUES ($1) RETURNING id"

func (p *Postgres) AddDefaultMedia(ctx context.Context, mediaID int) (int, error) {
	err := p.db.QueryRow(addDefaultMedia, mediaID).Scan(&mediaID)
	if err != nil {
		return 0, wrapPostgresError(err.(pgx.PgError).Code, err)
	}
	return mediaID, nil
}
