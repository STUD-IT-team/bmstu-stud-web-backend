package postgres

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

const getClub = `SELECT 
 name,
 short_name,
 description,
 type,
 logo,
 vk_url, 
 tg_url
 FROM club WHERE id = $1
`

func (pgs *Postgres) GetClub(id int) (*domain.Club, error) {
	c := domain.Club{}
	err := pgs.db.QueryRow(getClub, id).Scan(
		&c.Name,
		&c.ShortName,
		&c.Description,
		&c.Type,
		&c.LogoId,
		&c.VkUrl,
		&c.TgUrl,
	)
	c.ID = id

	if err != nil {
		return nil, err
	}
	return &c, nil
}

const getAllClub = `SELECT
 id,
 name,
 short_name,
 description,
 type,
 logo,
 vk_url, 
 tg_url
FROM club
`

func (s *Postgres) GetAllClub() ([]domain.Club, error) {
	carr := []domain.Club{}
	rows, err := s.db.Query(getAllClub)

	if err != nil {
		return []domain.Club{}, err
	}

	for rows.Next() {
		var c domain.Club
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.ShortName,
			&c.Description,
			&c.Type,
			&c.LogoId,
			&c.VkUrl,
			&c.TgUrl,
		)
		if err != nil {
			return []domain.Club{}, err
		}
		carr = append(carr, c)
	}

	return carr, nil
}

const getClubOrgs = `
SELECT
	role_name,
	role_spec,
	mem.id,
	mem.hash_password,
	mem.login,
	mem.media_id,
	mem.telegram,
	mem.vk,
	mem.name,
	mem.role_id,
	mem.is_admin,
	clubs.name as club_name
FROM club_org
JOIN
(
	SELECT
		id,
		hash_password,
		login,
		media_id,
		telegram,
		vk,
		name,
		role_id,
		is_admin
		FROM member
) mem
ON (mem.id = club_org.member_id)
JOIN 
(
	SELECT
	    id,
        name
    FROM club
) as clubs
ON (club_org.club_id = clubs.id)
WHERE club_id = $1
`

func (s *Postgres) GetClubOrgs(clubID int) ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getClubOrgs, clubID)
	if err != nil {
		return []domain.ClubOrg{}, err
	}
	for rows.Next() {
		c := domain.ClubOrg{}
		err = rows.Scan(
			&c.RoleName,
			&c.RoleSpec,
			&c.ID,
			&c.HashPassword,
			&c.Login,
			&c.MediaID,
			&c.Telegram,
			&c.Vk,
			&c.Name,
			&c.RoleID,
			&c.IsAdmin,
			&c.ClubName,
		)
		if err != nil {
			return []domain.ClubOrg{}, err
		}
		oarr = append(oarr, c)
	}
	return oarr, nil
}

const getClubSubOrgs = `
SELECT
	role_name,
	role_spec,
	mem.id,
	mem.hash_password,
	mem.login,
	mem.media_id,
	mem.telegram,
	mem.vk,
	mem.name,
	mem.role_id,
	mem.is_admin,
	clubs.name as club_name
FROM club_org
JOIN
(
	SELECT
		id,
		hash_password,
		login,
		media_id,
		telegram,
		vk,
		name,
		role_id,
		is_admin
		FROM member
) mem
ON (mem.id = club_org.member_id)
JOIN 
(
	SELECT
	    id,
        name
    FROM club
) as clubs
ON (club_org.club_id = clubs.id)
WHERE club_id = ANY((SELECT id FROM club WHERE parent_id = $1))
`

func (s *Postgres) GetClubSubOrgs(clubID int) ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getClubSubOrgs, clubID)
	if err != nil {
		return []domain.ClubOrg{}, err
	}
	for rows.Next() {
		c := domain.ClubOrg{}
		err = rows.Scan(
			&c.RoleName,
			&c.RoleSpec,
			&c.ID,
			&c.HashPassword,
			&c.Login,
			&c.MediaID,
			&c.Telegram,
			&c.Vk,
			&c.Name,
			&c.RoleID,
			&c.IsAdmin,
			&c.ClubName,
		)
		if err != nil {
			return []domain.ClubOrg{}, err
		}
		c.ClubID = clubID
		oarr = append(oarr, c)
	}
	return oarr, nil
}

const getAllClubOrgs = `
SELECT
	role_name,
	role_spec,
	member_id,
	club_id,
	mem.hash_password,
	mem.login,
	mem.media_id,
	mem.telegram,
	mem.vk,
	mem.name,
	mem.role_id,
	mem.is_admin,
	clubs.name as club_name
FROM club_org
JOIN
(
	SELECT
		id,
		hash_password,
		login,
		media_id,
		telegram,
		vk,
		name,
		role_id,
		is_admin
		FROM member
) mem
ON (mem.id = club_org.member_id)
JOIN 
(
	SELECT
	    id,
        name
    FROM club
) as clubs
ON (club_org.club_id = clubs.id)
`

func (s *Postgres) GetAllClubOrgs() ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getAllClubOrgs)
	if err != nil {
		return []domain.ClubOrg{}, err
	}
	for rows.Next() {
		c := domain.ClubOrg{}
		err = rows.Scan(
			&c.RoleName,
			&c.RoleSpec,
			&c.ID,
			&c.ClubID,
			&c.HashPassword,
			&c.Login,
			&c.MediaID,
			&c.Telegram,
			&c.Vk,
			&c.Name,
			&c.RoleID,
			&c.IsAdmin,
			&c.ClubName,
		)
		if err != nil {
			return []domain.ClubOrg{}, err
		}
		oarr = append(oarr, c)
	}
	return oarr, nil
}
