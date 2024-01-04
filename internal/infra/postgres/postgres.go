package postgres

import (
	"context"
	"fmt"
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

func (p *Postgres) GetAllFeed(ctx context.Context) ([]domain.Feed, error) {
	var feeds []domain.Feed

	query := "SELECT id, description FROM events"

	rows, err := p.db.Query(query)
	if err != nil {
		return []domain.Feed{}, err
	}

	for rows.Next() {
		var feed domain.Feed
		rows.Scan(&feed.ID, &feed.Description)
		feeds = append(feeds, feed)
	}

	return feeds, nil
}

func (p *Postgres) GetFeed(ctx context.Context, id int) (domain.Feed, error) {
	var feed domain.Feed

	query := fmt.Sprintf("SELECT id, description, reg_url FROM events WHERE id=%d", id)

	err := p.db.QueryRow(query).Scan(&feed.ID, &feed.Description, &feed.RegistationURL)
	if err != nil {
		return domain.Feed{}, err
	}

	return feed, nil
}
