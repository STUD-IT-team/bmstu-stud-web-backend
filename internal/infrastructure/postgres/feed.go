package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllFeedQuery = "SELECT id, title, description FROM events"

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

const getFeedQuery = "SELECT id, title, description, reg_url FROM events WHERE id=$1"

func (p *Postgres) GetFeed(_ context.Context, id int) (domain.Feed, error) {
	var feed domain.Feed

	err := p.db.QueryRow(getFeedQuery, id).Scan(&feed.ID, &feed.Title, &feed.Description, &feed.RegistrationURL)
	if err != nil {
		return domain.Feed{}, err
	}

	return feed, nil
}

const deleteFeedQuery = "DELETE FROM events WHERE id=$1"

func (p *Postgres) DeleteFeed(_ context.Context, id int) error {
	_, err := p.db.Exec(deleteFeedQuery, id)
	if err != nil {
		return err
	}

	return nil
}

const putFeedQuery = "UPDATE events SET title=$1, description=$2, reg_url=$3, created_by=$4, date=$5 WHERE id=$6"

func (p *Postgres) UpdateFeed(_ context.Context, feed domain.Feed) error {
	_, err := p.db.Exec(putFeedQuery,
		feed.Title,
		feed.Description,
		feed.RegistrationURL,
		feed.CreatedBy,
		feed.UpdatedAt,
		feed.ID,
	)
	if err != nil {
		return fmt.Errorf("can't update feed on postgres %w", err)
	}

	return nil
}

const getFeedByFilterLimitAndOffsetQuery = `SELECT id, title, description, reg_url, created_at, created_by 
											FROM events ORDER BY id LIMIT $1 OFFSET $2`

func (p *Postgres) GetFeedByFilterLimitAndOffset(_ context.Context, limit, offset int) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getFeedByFilterLimitAndOffsetQuery, limit, offset)
	if err != nil {
		return []domain.Feed{}, err
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Description, &feed.RegistrationURL, &feed.CreatedAt, &feed.CreatedBy)
		if err != nil {
			return []domain.Feed{}, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

const getFeedByFilterIdLastAndOffsetQuery = `SELECT id, title, description, reg_url, created_at, created_by 
											FROM events  WHERE id > $1 ORDER BY id LIMIT $2`

func (p *Postgres) GetFeedByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.Feed, error) {
	var feeds []domain.Feed

	rows, err := p.db.Query(getFeedByFilterIdLastAndOffsetQuery, idLast, offset)
	if err != nil {
		return []domain.Feed{}, err
	}

	for rows.Next() {
		var feed domain.Feed

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Description, &feed.RegistrationURL, &feed.CreatedAt, &feed.CreatedBy)
		if err != nil {
			return []domain.Feed{}, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
