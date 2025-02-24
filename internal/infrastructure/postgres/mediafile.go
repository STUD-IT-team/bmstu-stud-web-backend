package postgres

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getMediaFile = "SELECT name, key FROM mediafile WHERE id = $1 AND id > 0"

func (p *Postgres) GetMediaFile(id int) (*domain.MediaFile, error) {
	f := domain.MediaFile{}
	err := p.db.QueryRow(getMediaFile, id).Scan(&f.Name, &f.Key)
	if err != nil {
		return nil, wrapPostgresError(err)
	}
	f.ID = id
	return &f, nil
}

const getMediaFiles = "SELECT id, name, key FROM mediafile WHERE id = ANY($1) AND id > 0"

func (p *Postgres) GetMediaFiles(ids []int) (map[int]domain.MediaFile, error) {
	m := make(map[int]domain.MediaFile)
	rows, err := p.db.Query(getMediaFiles, ids)
	if err != nil {
		return nil, wrapPostgresError(err)
	}
	for rows.Next() {
		media := domain.MediaFile{}
		err := rows.Scan(&media.ID, &media.Name, &media.Key)
		if err != nil {
			return nil, wrapPostgresError(err)
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
		return 0, wrapPostgresError(err)
	}
	return id, nil
}

const updateMediaFile = "UPDATE mediafile SET name = $1, key = $2 WHERE id = $3"

func (p *Postgres) UpdateMediaFile(id int, name, key string) error {
	tag, err := p.db.Exec(updateMediaFile, name, key, id)
	if tag.RowsAffected() == 0 {
		return wrapPostgresError(err)
	}
	if err != nil {
		return wrapPostgresError(err)
	}
	return nil
}

const deleteMediaFile = "DELETE FROM mediafile WHERE id = $1 AND id > 0"

func (p *Postgres) DeleteMediaFile(id int) error {
	_, err := p.db.Exec(deleteMediaFile, id)
	if err != nil {
		return wrapPostgresError(err)
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
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		med := domain.MediaFile{}
		err := rows.Scan(&med.ID, &med.Name, &med.Key)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		res = append(res, med)
	}
	return res, nil
}

const deleteMediaFiles = "DELETE FROM mediafile WHERE key = ANY($1) AND id > 0"

func (p *Postgres) DeleteMediaFiles(ctx context.Context, keys []string) error {
	_, err := p.db.Exec(deleteMediaFiles, keys)
	if err != nil {
		return wrapPostgresError(err)
	}
	return nil
}

const getAllMediaKeys = "SELECT key FROM mediafile WHERE id > 0"

func (p *Postgres) GetAllMediaKeys(ctx context.Context) ([]string, error) {
	rows, err := p.db.Query(getAllMediaKeys)
	if err != nil {
		return nil, wrapPostgresError(err)
	}
	defer rows.Close()

	var keys []string
	for rows.Next() {
		var key string
		err := rows.Scan(&key)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		keys = append(keys, key)
	}
	return keys, nil
}

const getDefautlMedia = "SELECT id, media_id FROM default_media WHERE id = $1"

func (p *Postgres) GetDefautlMedia(ctx context.Context, id int) (*domain.DefaultMedia, error) {
	var d domain.DefaultMedia
	err := p.db.QueryRow(getDefautlMedia, id).Scan(&d.ID, &d.MediaID)
	if err != nil {
		return nil, wrapPostgresError(err)
	}
	return &d, nil
}

const getAllDefaultMedia = "SELECT id, media_id FROM default_media"

func (p *Postgres) GetAllDefaultMedia(ctx context.Context) ([]domain.DefaultMedia, error) {
	defaultMedia := make([]domain.DefaultMedia, 0)
	rows, err := p.db.Query(getAllDefaultMedia)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		d := domain.DefaultMedia{}
		err := rows.Scan(&d.ID, &d.MediaID)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		defaultMedia = append(defaultMedia, d)
	}

	if len(defaultMedia) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return defaultMedia, nil
}

const addDefaultMedia = "INSERT INTO default_media (media_id) VALUES ($1) RETURNING id"

func (p *Postgres) AddDefaultMedia(ctx context.Context, mediaID int) (int, error) {
	err := p.db.QueryRow(addDefaultMedia, mediaID).Scan(&mediaID)
	if err != nil {
		return 0, wrapPostgresError(err)
	}
	return mediaID, nil
}

const deleteDefaultMedia = "DELETE FROM default_media WHERE id = $1"

func (p *Postgres) DeleteDefaultMedia(ctx context.Context, id int) error {
	tag, err := p.db.Exec(deleteDefaultMedia, id)
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}
	if err != nil {
		return wrapPostgresError(err)
	}
	return nil
}

const updateDefaultMedia = "UPDATE default_media SET media_id = $1 WHERE id = $2"

func (p *Postgres) UpdateDefaultMedia(ctx context.Context, id, mediaID int) error {
	res, err := p.db.Exec(updateDefaultMedia, mediaID, id)
	if res.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}
	if err != nil {
		return wrapPostgresError(err)
	}
	return nil
}

const getMediaFileByKey = "SELECT id, name, key FROM mediafile WHERE key = $1"

func (p *Postgres) GetMediaFileByKey(ctx context.Context, key string) (*domain.MediaFile, error) {
	f := domain.MediaFile{}
	err := p.db.QueryRow(getMediaFileByKey, key).Scan(&f.ID, &f.Name, &f.Key)
	if err != nil {
		return nil, wrapPostgresError(err)
	}
	return &f, nil
}

const getRandomDefaultMedia = "SELECT id, media_id FROM default_media OFFSET FLOOR(RANDOM() * (SELECT COUNT(*) FROM default_media)) LIMIT 1"

func (p *Postgres) GetRandomDefaultMedia(ctx context.Context) (*domain.DefaultMedia, error) {
	var d domain.DefaultMedia
	err := p.db.QueryRow(getRandomDefaultMedia).Scan(&d.ID, &d.MediaID)
	if err != nil {
		return nil, wrapPostgresError(err)
	}
	return &d, nil
}

const getActiveMainVideo = `
SELECT id, name, key, club_id, current FROM main_video
WHERE current = true AND club_id = 0
`

func (p *Postgres) GetActiveMainVideo(ctx context.Context) (*domain.MainVideo, error) {
	var m domain.MainVideo
	err := p.db.QueryRow(getActiveMainVideo).Scan(&m.ID, &m.Name, &m.Key, &m.ClubID, &m.Current)
	if err != nil {
		return nil, wrapPostgresError(err)
	}
	return &m, nil
}
