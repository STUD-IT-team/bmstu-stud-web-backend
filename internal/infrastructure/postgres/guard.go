package postgres

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const addMember = `INSERT INTO member (login, hash_password, name, telegram, vk, media_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

func (p *Postgres) AddMember(ctx context.Context, mem *domain.Member) (int, error) {
	var id int
	err := p.db.QueryRow(addMember, mem.Login, mem.HashPassword, mem.Name, mem.Telegram, mem.Vk, mem.MediaID).Scan(&id)
	if err != nil {
		return 0, wrapPostgresError(err)
	}
	return id, nil
}

const getMemberHash = `SELECT hash_password FROM member WHERE login = $1`

func (p *Postgres) GetMemberHash(_ context.Context, login string) (string, error) {
	var hash string
	err := p.db.QueryRow(getMemberHash, login).Scan(&hash)
	if err != nil {
		return "", wrapPostgresError(err)
	}
	return hash, nil
}
