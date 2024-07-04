package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllMembersQuery = "SELECT id, hash_password, login, media_id, telegram, vk, name, role_id, is_admin FROM member"

func (p *Postgres) GetAllMembers(_ context.Context) ([]domain.Member, error) {
	var members []domain.Member

	rows, err := p.db.Query(getAllMembersQuery)
	if err != nil {
		return []domain.Member{}, err
	}

	for rows.Next() {
		var member domain.Member

		err = rows.Scan(
			&member.ID,
			&member.HashPassword,
			&member.Login,
			&member.MediaID,
			&member.Telegram,
			&member.Vk,
			&member.Name,
			&member.RoleID,
			&member.IsAdmin,
		)

		if err != nil {
			return []domain.Member{}, err
		}

		members = append(members, member)
	}

	if len(members) == 0 {
		return []domain.Member{}, fmt.Errorf("no members found")
	}

	return members, nil
}

const getMemberQuery = "SELECT id, hash_password, login, media_id, telegram, vk, name, role_id, is_admin FROM member WHERE id=$1"

func (p *Postgres) GetMember(ctx context.Context, id int) (domain.Member, error) {
	var member domain.Member

	err := p.db.QueryRow(
		getMemberQuery,
		id,
	).Scan(
		&member.ID,
		&member.HashPassword,
		&member.Login,
		&member.MediaID,
		&member.Telegram,
		&member.Vk,
		&member.Name,
		&member.RoleID,
		&member.IsAdmin,
	)

	if err != nil {
		return domain.Member{}, err
	}

	return member, nil
}
