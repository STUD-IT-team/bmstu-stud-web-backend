package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllFeedQuery = "SELECT id, title, description FROM feed"

func (p *Postgres) GetAllFeed(_ context.Context) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getAllFeedQuery)
	if err != nil {
		return []domain.Feed{}, err
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Description)

		if err != nil {
			return []domain.Feed{}, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

const getFeedQuery = "SELECT id, title, description FROM feed WHERE id=$1"

func (p *Postgres) GetFeed(_ context.Context, id int) (domain.Feed, error) {
	var feed domain.Feed

	err := p.db.QueryRow(getFeedQuery, id).Scan(&feed.ID, &feed.Title, &feed.Description)
	if err != nil {
		return domain.Feed{}, err
	}

	return feed, nil
}

const getFeedEncountersQuery = "SELECT id, count, description, club_id FROM encounter WHERE club_id=$1"

func (p *Postgres) GetFeedEncounters(_ context.Context, id int) ([]domain.Encounter, error) {
	var encs []domain.Encounter

	rows, err := p.db.Query(getFeedEncountersQuery, id)
	if err != nil {
		return []domain.Encounter{}, err
	}

	for rows.Next() {
		var enc domain.Encounter

		err = rows.Scan(&enc.ClubID, &enc.Count, &enc.Description, &enc.ClubID)

		if err != nil {
			return []domain.Encounter{}, err
		}

		encs = append(encs, enc)
	}

	return encs, nil
}

const getFeedByTitleQuery = "SELECT id, title, description FROM feed WHERE title ILIKE $1"

func (p *Postgres) GetFeedByTitle(_ context.Context, title string) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getFeedByTitleQuery, "%"+title+"%")
	if err != nil {
		return []domain.Feed{}, err
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Description)

		if err != nil {
			return []domain.Feed{}, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

const deleteFeedQuery = "DELETE FROM feed WHERE id=$1"

func (p *Postgres) DeleteFeed(_ context.Context, id int) error {
	_, err := p.db.Exec(deleteFeedQuery, id)
	if err != nil {
		return err
	}

	return nil
}

const putFeedQuery = "UPDATE feed SET title=$1, description=$2, created_by=$4, date=$5 WHERE id=$6"

func (p *Postgres) UpdateFeed(_ context.Context, feed domain.Feed) error {
	_, err := p.db.Exec(putFeedQuery,
		feed.Title,
		feed.Description,
		feed.CreatedBy,
		feed.UpdatedAt,
		feed.ID,
	)
	if err != nil {
		return fmt.Errorf("can't update feed on postgres %w", err)
	}

	return nil
}

const getFeedByFilterLimitAndOffsetQuery = `SELECT id, title, description, created_at, created_by 
											FROM feed ORDER BY id LIMIT $1 OFFSET $2`

func (p *Postgres) GetFeedByFilterLimitAndOffset(_ context.Context, limit, offset int) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getFeedByFilterLimitAndOffsetQuery, limit, offset)
	if err != nil {
		return []domain.Feed{}, err
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Description, &feed.CreatedAt, &feed.CreatedBy)
		if err != nil {
			return []domain.Feed{}, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

const getFeedByFilterIdLastAndOffsetQuery = `SELECT id, title, description, created_at, created_by 
											FROM feed  WHERE id > $1 ORDER BY id LIMIT $2`

func (p *Postgres) GetFeedByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getFeedByFilterIdLastAndOffsetQuery, idLast, offset)
	if err != nil {
		return []domain.Feed{}, err
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Description, &feed.CreatedAt, &feed.CreatedBy)
		if err != nil {
			return []domain.Feed{}, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
