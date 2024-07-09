package postgres

import (
	"context"
	"database/sql"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/jackc/pgx"
)

const getClub = `SELECT 
 name,
 short_name,
 description,
 type,
 logo,
 parent_id,
 vk_url, 
 tg_url
 FROM club WHERE id = $1
`

const NoParentClubID = 0

func (pgs *Postgres) GetClub(_ context.Context, id int) (*domain.Club, error) {
	parentID := sql.NullInt64{}
	c := domain.Club{}
	err := pgs.db.QueryRow(getClub, id).Scan(
		&c.Name,
		&c.ShortName,
		&c.Description,
		&c.Type,
		&c.LogoId,
		&parentID,
		&c.VkUrl,
		&c.TgUrl,
	)
	if err != nil {
		return nil, err
	}
	c.ID = id
	if parentID.Valid {
		c.ParentID = int(parentID.Int64)
	} else {
		c.ParentID = NoParentClubID
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
 parent_id,
 vk_url, 
 tg_url
FROM club
`

func (s *Postgres) GetAllClub(_ context.Context) ([]domain.Club, error) {
	carr := []domain.Club{}
	rows, err := s.db.Query(getAllClub)

	if err != nil {
		return []domain.Club{}, err
	}

	for rows.Next() {
		var c domain.Club
		parentID := sql.NullInt64{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.ShortName,
			&c.Description,
			&c.Type,
			&c.LogoId,
			&parentID,
			&c.VkUrl,
			&c.TgUrl,
		)
		if err != nil {
			return []domain.Club{}, err
		}
		if parentID.Valid {
			c.ParentID = int(parentID.Int64)
		} else {
			c.ParentID = NoParentClubID
		}
		carr = append(carr, c)
	}

	return carr, nil
}

const getClubsByName = `SELECT id, name, short_name, description, type, logo, parent_id, vk_url, tg_url FROM club WHERE name ILIKE $1`

func (s *Postgres) GetClubsByName(_ context.Context, name string) ([]domain.Club, error) {
	carr := []domain.Club{}
	rows, err := s.db.Query(getClubsByName, "%"+name+"%")

	if err != nil {
		return []domain.Club{}, err
	}

	for rows.Next() {
		var c domain.Club
		parentID := sql.NullInt64{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.ShortName,
			&c.Description,
			&c.Type,
			&c.LogoId,
			&parentID,
			&c.VkUrl,
			&c.TgUrl,
		)
		if err != nil {
			return []domain.Club{}, err
		}
		if parentID.Valid {
			c.ParentID = int(parentID.Int64)
		} else {
			c.ParentID = NoParentClubID
		}
		carr = append(carr, c)

	}
	return carr, nil
}

const getClubsByType = `SELECT id, name, short_name, description, type, logo, parent_id, vk_url, tg_url FROM club WHERE type ILIKE $1`

func (s *Postgres) GetClubsByType(_ context.Context, type_ string) ([]domain.Club, error) {
	carr := []domain.Club{}
	rows, err := s.db.Query(getClubsByType, "%"+type_+"%")

	if err != nil {
		return []domain.Club{}, err
	}

	for rows.Next() {
		var c domain.Club
		parentID := sql.NullInt64{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.ShortName,
			&c.Description,
			&c.Type,
			&c.LogoId,
			&parentID,
			&c.VkUrl,
			&c.TgUrl,
		)
		if err != nil {
			return []domain.Club{}, err
		}
		if parentID.Valid {
			c.ParentID = int(parentID.Int64)
		} else {
			c.ParentID = NoParentClubID
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

func (s *Postgres) GetClubOrgs(_ context.Context, clubID int) ([]domain.ClubOrg, error) {
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

const getClubsOrgs = `
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
	clubs.name as club_name,
	clubs.id as club_id
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
WHERE club_id = ANY($1)
`

func (s *Postgres) GetClubsOrgs(_ context.Context, clubIDs []int) ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getClubsOrgs, clubIDs)
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
			&c.ClubID,
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

func (s *Postgres) GetClubSubOrgs(_ context.Context, clubID int) ([]domain.ClubOrg, error) {
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

func (s *Postgres) GetAllClubOrgs(_ context.Context) ([]domain.ClubOrg, error) {
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

const addClub = `
INSERT INTO club (
    name,
    short_name,
    description,
    type,
	logo,
    parent_id,
    vk_url, 
    tg_url
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id
`

func (s *Postgres) AddClub(_ context.Context, c *domain.Club) (int, error) {
	row := s.db.QueryRow(addClub,
		c.Name,
		c.ShortName,
		c.Description,
		c.Type,
		c.LogoId,
		c.ParentID,
		c.VkUrl,
		c.TgUrl,
	)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, wrapPostgresError(err.(pgx.PgError).Code, err)
	}
	return id, nil
}

const addOrgs = `
INSERT INTO club_org (
	role_name,
    role_spec,
    member_id,
    club_id
) VALUES ($1, $2, $3, $4)
`

func (s *Postgres) AddOrgs(_ context.Context, orgs []domain.ClubOrg) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, org := range orgs {
		_, err = tx.Exec(addOrgs, org.RoleName, org.RoleSpec, org.ID, org.ClubID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

const getClubMediaFiles = `
SELECT
    club_photo.id,
	club_photo.ref_num,
	photo.name,
	photo.image_url
FROM club_photo
JOIN
(
	SELECT
	    name,
		image_url,
		id
	FROM mediafile
) as photo
ON (club_photo.media_id = photo.id)
WHERE club_id = $1
`

func (s *Postgres) GetClubMediaFiles(clubID int) ([]domain.ClubPhoto, error) {
	ph := []domain.ClubPhoto{}
	rows, err := s.db.Query(getClubMediaFiles, clubID)
	if err != nil {
		return []domain.ClubPhoto{}, err
	}

	for rows.Next() {
		p := domain.ClubPhoto{}
		err := rows.Scan(&p.ID, &p.RefNumber, &p.Name, &p.ImageUrl)
		if err != nil {
			return []domain.ClubPhoto{}, err
		}
		p.ClubID = clubID
		ph = append(ph, p)
	}
	return ph, nil
}

const deleteClub = "DELETE FROM club WHERE id = $1"
const deleteClubOrgs = "DELETE FROM club_org WHERE club_id = $1"
const deleteClubPhotos = "DELETE FROM club_photo WHERE club_id = $1"
const deleteClubEncounters = "DELETE FROM encounter WHERE club_id = $1"
const updateClubParents = "UPDATE club SET parent_id=null WHERE parent_id = $1"

func (s *Postgres) DeleteClubWithOrgs(_ context.Context, clubID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(deleteClub, clubID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(deleteClubOrgs, clubID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(deleteClubPhotos, clubID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(deleteClubEncounters, clubID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(updateClubParents, clubID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

const updateClub = `
UPDATE club
SET name=$1,
    short_name=$2,
    description=$3,
    type=$4,
	logo=$5,
    parent_id=$6,
    vk_url=$7, 
    tg_url=$8
WHERE id = $9
`

func (s *Postgres) UpdateClub(_ context.Context, c *domain.Club, o []domain.ClubOrg) error {
	tx, err := s.db.Begin()
	if err != nil {
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}

	_, err = tx.Exec(updateClub,
		c.Name,
		c.ShortName,
		c.Description,
		c.Type,
		c.LogoId,
		c.ParentID,
		c.VkUrl,
		c.TgUrl,
		c.ID,
	)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}

	_, err = tx.Exec(deleteClubOrgs, c.ID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}

	for _, org := range o {
		_, err = tx.Exec(addOrgs, org.RoleName, org.RoleSpec, org.ID, c.ID)
		if err != nil {
			tx.Rollback()
			return wrapPostgresError(err.(pgx.PgError).Code, err)
		}
	}
	err = tx.Commit()
	return wrapPostgresError(err.(pgx.PgError).Code, err)
}
