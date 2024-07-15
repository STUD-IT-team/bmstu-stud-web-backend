package postgres

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllFeedQuery = "SELECT id, title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by FROM feed"

func (p *Postgres) GetAllFeed(_ context.Context) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getAllFeedQuery)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Approved,
			&feed.Description, &feed.MediaID, &feed.VkPostUrl,
			&feed.UpdatedAt, &feed.CreatedAt, &feed.Views, &feed.CreatedBy)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		feeds = append(feeds, feed)
	}

	if len(feeds) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return feeds, nil
}

const getFeedQuery = "SELECT id, title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by FROM feed WHERE id=$1"

func (p *Postgres) GetFeed(_ context.Context, id int) (*domain.Feed, error) {
	var feed domain.Feed

	err := p.db.QueryRow(getFeedQuery, id).Scan(&feed.ID, &feed.Title, &feed.Approved,
		&feed.Description, &feed.MediaID, &feed.VkPostUrl,
		&feed.UpdatedAt, &feed.CreatedAt, &feed.Views, &feed.CreatedBy)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	return &feed, nil
}

const getFeedEncountersQuery = "SELECT id, count, description, club_id FROM encounter WHERE club_id=$1"

func (p *Postgres) GetFeedEncounters(_ context.Context, id int) ([]domain.Encounter, error) {
	var encs []domain.Encounter

	rows, err := p.db.Query(getFeedEncountersQuery, id)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var enc domain.Encounter

		err = rows.Scan(&enc.ID, &enc.Count, &enc.Description, &enc.ClubID)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		encs = append(encs, enc)
	}

	if len(encs) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return encs, nil
}

const getFeedByTitleQuery = "SELECT id, title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by FROM feed WHERE title ILIKE $1"

func (p *Postgres) GetFeedByTitle(_ context.Context, title string) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getFeedByTitleQuery, "%"+title+"%")
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Approved,
			&feed.Description, &feed.MediaID, &feed.VkPostUrl,
			&feed.UpdatedAt, &feed.CreatedAt, &feed.Views, &feed.CreatedBy)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		feeds = append(feeds, feed)
	}

	if len(feeds) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return feeds, nil
}

const postFeedQuery = `INSERT INTO feed (title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

func (p *Postgres) PostFeed(_ context.Context, feed *domain.Feed) error {
	_, err := p.db.Exec(postFeedQuery,
		feed.Title,
		feed.Approved,
		feed.Description,
		feed.MediaID,
		feed.VkPostUrl,
		feed.UpdatedAt,
		feed.CreatedAt,
		feed.Views,
		feed.CreatedBy,
	)

	if err != nil {
		return wrapPostgresError(err)
	}

	return nil
}

const deleteFeedQuery = "DELETE FROM feed WHERE id=$1"

func (p *Postgres) DeleteFeed(_ context.Context, id int) error {
	tag, err := p.db.Exec(deleteFeedQuery, id)
	if err != nil {
		return wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}

	return nil
}

const updateFeedQuery = `
UPDATE feed SET 
title=$1, 
approved=$2, 
description=$3, 
media_id=$4, 
vk_post_url=$5, 
updated_at=$6, 
created_at=$7, 
views=$8, 
created_by=$9 WHERE id=$10`

func (p *Postgres) UpdateFeed(_ context.Context, feed *domain.Feed) error {
	tag, err := p.db.Exec(updateFeedQuery,
		feed.Title,
		feed.Approved,
		feed.Description,
		feed.MediaID,
		feed.VkPostUrl,
		feed.UpdatedAt,
		feed.CreatedAt,
		feed.Views,
		feed.CreatedBy,
		feed.ID,
	)
	if err != nil {
		return wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}

	return nil
}

const postEncounterquery = `INSERT INTO encounter (count, description, club_id) VALUES ($1, $2, $3)`

func (p *Postgres) PostEncounter(_ context.Context, enc *domain.Encounter) error {
	_, err := p.db.Exec(postEncounterquery, enc.Count, enc.Description, enc.ClubID)
	if err != nil {
		return wrapPostgresError(err)
	}

	return nil
}

const deleteEncounterQuery = "DELETE FROM encounter WHERE id=$1"

func (p *Postgres) DeleteEncounter(_ context.Context, id int) error {
	tag, err := p.db.Exec(deleteEncounterQuery, id)
	if err != nil {
		return wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}

	return nil
}

const updateEncounterQuery = `UPDATE encounter SET count=$1, description=$2, club_id=$3 WHERE id=$4`

func (p *Postgres) UpdateEncounter(_ context.Context, enc *domain.Encounter) error {
	tag, err := p.db.Exec(updateEncounterQuery, enc.Count, enc.Description, enc.ClubID, enc.ID)
	if err != nil {
		return wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}

	return nil
}
