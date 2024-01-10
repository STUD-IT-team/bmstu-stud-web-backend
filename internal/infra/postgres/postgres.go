package postgres

import (
	"context"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

	"github.com/jackc/pgx"
)

type Postgres struct {
	db *pgx.ConnPool
}

func NewPostgres(databaseURL string) (*Postgres, error) {
	connConf, err := pgx.ParseURI(databaseURL)
	if err != nil {
		return nil, err
	}

	conf := pgx.ConnPoolConfig{ConnConfig: connConf, MaxConnections: 10, AcquireTimeout: time.Second * 1}
	db, err := pgx.NewConnPool(conf)
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

const getAllFeedQuery = "SELECT id, title, description FROM events"

func (p *Postgres) GetAllFeed(ctx context.Context) ([]domain.Feed, error) {
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

func (p *Postgres) GetFeed(ctx context.Context, id int) (domain.Feed, error) {
	var feed domain.Feed

	err := p.db.QueryRow(getFeedQuery, id).Scan(&feed.ID, &feed.Title, &feed.Description, &feed.RegistationURL)
	if err != nil {
		return domain.Feed{}, err
	}

	return feed, nil
}
