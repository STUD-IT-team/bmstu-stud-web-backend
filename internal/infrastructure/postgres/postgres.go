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

const getMemberByLoginQuery = "SELECT id, login, password FROM members WHERE login=$1;"

func (p *Postgres) GetMemberByLogin(_ context.Context, login string) (domain.Member, error) {
	const op = "postgres.GetUserByLogin"

	var user domain.Member

	err := p.db.QueryRow(getMemberByLoginQuery, login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Member{}, fmt.Errorf("%s: %w", op, domain.ErrNotFound)
		}

		return domain.Member{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
