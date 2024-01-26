package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

	"github.com/jackc/pgx"
)

type Postgres struct {
	db *pgx.ConnPool
}

const maxConn = 10

func NewPostgres(databaseURL string) (*Postgres, error) {
	connConf, err := pgx.ParseURI(databaseURL)
	if err != nil {
		return nil, err
	}

	conf := pgx.ConnPoolConfig{ConnConfig: connConf, MaxConnections: maxConn, AcquireTimeout: time.Second * 1}
	db, err := pgx.NewConnPool(conf)

	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

const getAllFeedQuery = "SELECT ID, TITLE, DESCRIPTION FROM EVENTS"

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

const getFeedQuery = "SELECT ID, TITLE, DESCRIPTION, REG_URL FROM EVENTS WHERE ID=$1"

func (p *Postgres) GetFeed(_ context.Context, id int) (domain.Feed, error) {
	var feed domain.Feed

	err := p.db.QueryRow(getFeedQuery, id).Scan(&feed.ID, &feed.Title, &feed.Description, &feed.RegistrationURL)
	if err != nil {
		return domain.Feed{}, err
	}

	return feed, nil
}

const getUserIDQuery = "SELECT id FROM users WHERE email =$1  AND hash_password = crypt($2, hash_password);"

func (p *Postgres) GetUserID(_ context.Context, email string, password string) (string, error) {
	const op = "postgres.GetUserID"

	var userID string
	err := p.db.QueryRow(getUserIDQuery, email, password).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, domain.ErrUserNotFound)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return userID, nil
}
