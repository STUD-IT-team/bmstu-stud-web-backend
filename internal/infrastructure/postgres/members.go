package postgres

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllMembersQuery = "SELECT id, login, media_id, telegram, vk, name, is_admin FROM member"

func (p *Postgres) GetAllMembers(_ context.Context) ([]domain.Member, error) {
	var members []domain.Member

	rows, err := p.db.Query(getAllMembersQuery)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var member domain.Member

		err = rows.Scan(
			&member.ID,
			&member.Login,
			&member.MediaID,
			&member.Telegram,
			&member.Vk,
			&member.Name,
			&member.IsAdmin,
		)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		members = append(members, member)
	}

	if len(members) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return members, nil
}

const getMemberQuery = "SELECT id, login, media_id, telegram, vk, name, is_admin FROM member WHERE id=$1"

func (p *Postgres) GetMember(ctx context.Context, id int) (*domain.Member, error) {
	var member domain.Member

	err := p.db.QueryRow(
		getMemberQuery,
		id,
	).Scan(
		&member.ID,
		&member.Login,
		&member.MediaID,
		&member.Telegram,
		&member.Vk,
		&member.Name,
		&member.IsAdmin,
	)

	if err != nil {
		return nil, wrapPostgresError(err)
	}

	return &member, nil
}

const getMembersByNameQuery = "SELECT id, login, media_id, telegram, vk, name, is_admin FROM member WHERE name ILIKE $1"

func (p *Postgres) GetMembersByName(_ context.Context, name string) ([]domain.Member, error) {
	var members []domain.Member

	rows, err := p.db.Query(getMembersByNameQuery, "%"+name+"%")
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var member domain.Member

		err = rows.Scan(
			&member.ID,
			&member.Login,
			&member.MediaID,
			&member.Telegram,
			&member.Vk,
			&member.Name,
			&member.IsAdmin,
		)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		members = append(members, member)
	}

	if len(members) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return members, nil
}

const postMemberQuery = `INSERT INTO member 
	(login, media_id, telegram, vk, name, is_admin) 
	VALUES ($1, $2, $3, $4, $5, $6)`

func (p *Postgres) PostMember(ctx context.Context, member *domain.Member) error {
	_, err := p.db.Exec(
		postMemberQuery,
		member.Login,
		member.MediaID,
		member.Telegram,
		member.Vk,
		member.Name,
		member.IsAdmin,
	)

	if err != nil {
		return wrapPostgresError(err)
	}

	return nil
}

const deleteMemberQuery = "DELETE FROM member WHERE id=$1"

func (p *Postgres) DeleteMember(ctx context.Context, id int) error {
	tag, err := p.db.Exec(
		deleteMemberQuery,
		id,
	)
	if err != nil {
		return wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}

	return nil
}

const updateMemberQuery = `
UPDATE member SET
login=$1, 
media_id=$2, 
telegram=$3, 
vk=$4, 
name=$5, 
is_admin=$6
WHERE id=$7`

func (p *Postgres) UpdateMember(ctx context.Context, member *domain.Member) error {
	tag, err := p.db.Exec(
		updateMemberQuery,
		member.Login,
		member.MediaID,
		member.Telegram,
		member.Vk,
		member.Name,
		member.IsAdmin,
		member.ID,
	)
	if err != nil {
		return wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}

	return nil
}

const getMemberByLoginQuery = "SELECT id, login, hash_password, media_id, telegram, vk, name, is_admin FROM member WHERE login=$1;"

func (p *Postgres) GetMemberByLogin(_ context.Context, login string) (*domain.Member, error) {
	var user domain.Member

	err := p.db.QueryRow(getMemberByLoginQuery, login).Scan(&user.ID, &user.Login, &user.HashPassword, &user.MediaID, &user.Telegram, &user.Vk, &user.Name, &user.IsAdmin)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	return &user, nil
}
