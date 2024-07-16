package postgres

import (
	"context"
	"database/sql"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
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
 FROM club WHERE id = $1 AND id > 0
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
		return nil, wrapPostgresError(err)
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
WHERE id > 0
`

func (s *Postgres) GetAllClub(_ context.Context) ([]domain.Club, error) {
	carr := []domain.Club{}
	rows, err := s.db.Query(getAllClub)

	if err != nil {
		return nil, wrapPostgresError(err)
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
			return nil, wrapPostgresError(err)
		}
		if parentID.Valid {
			c.ParentID = int(parentID.Int64)
		} else {
			c.ParentID = NoParentClubID
		}
		carr = append(carr, c)
	}
	if len(carr) == 0 {
		return nil, ErrPostgresNotFoundError
	}
	return carr, nil
}

const getClubsByName = `SELECT id, name, short_name, description, type, logo, parent_id, vk_url, tg_url FROM club WHERE name ILIKE $1 AND id > 0`

func (s *Postgres) GetClubsByName(_ context.Context, name string) ([]domain.Club, error) {
	carr := []domain.Club{}
	rows, err := s.db.Query(getClubsByName, "%"+name+"%")

	if err != nil {
		return nil, wrapPostgresError(err)
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
			return nil, wrapPostgresError(err)
		}
		if parentID.Valid {
			c.ParentID = int(parentID.Int64)
		} else {
			c.ParentID = NoParentClubID
		}
		carr = append(carr, c)

	}
	if len(carr) == 0 {
		return nil, ErrPostgresNotFoundError
	}
	return carr, nil
}

const getClubsByType = `SELECT id, name, short_name, description, type, logo, parent_id, vk_url, tg_url FROM club WHERE type ILIKE $1 AND id > 0`

func (s *Postgres) GetClubsByType(_ context.Context, type_ string) ([]domain.Club, error) {
	carr := []domain.Club{}
	rows, err := s.db.Query(getClubsByType, "%"+type_+"%")

	if err != nil {
		return nil, wrapPostgresError(err)
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
			return nil, wrapPostgresError(err)
		}
		if parentID.Valid {
			c.ParentID = int(parentID.Int64)
		} else {
			c.ParentID = NoParentClubID
		}
		carr = append(carr, c)

	}
	if len(carr) == 0 {
		return nil, ErrPostgresNotFoundError
	}
	return carr, nil
}

const getClubOrgs = `
SELECT
    orgs.role_name,
    orgs.role_spec,
	mem.id,
	mem.hash_password,
	mem.login,
	mem.media_id,
	mem.telegram,
	mem.vk,
	mem.name,
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
JOIN
(
	SELECT
        id,
		role_name,
		role_spec
	FROM club_role
    WHERE role_clearance = 2
) as orgs
ON orgs.id = club_org.role_id
WHERE club_id = $1 AND club_id > 0
`

func (s *Postgres) GetClubOrgs(_ context.Context, clubID int) ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getClubOrgs, clubID)
	if err != nil {
		return nil, wrapPostgresError(err)
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
			&c.IsAdmin,
			&c.ClubName,
		)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		oarr = append(oarr, c)
	}
	if len(oarr) == 0 {
		return nil, ErrPostgresNotFoundError
	}
	return oarr, nil
}

const getClubsOrgs = `
SELECT
    orgs.role_name,
    orgs.role_spec,
	mem.id,
	mem.hash_password,
	mem.login,
	mem.media_id,
	mem.telegram,
	mem.vk,
	mem.name,
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
JOIN
(
	SELECT
        id,
		role_name,
		role_spec
	FROM club_role
    WHERE role_clearance = 2
) as orgs
ON orgs.id = club_org.role_id
WHERE club_id = ANY($1) AND club_id > 0
`

func (s *Postgres) GetClubsOrgs(_ context.Context, clubIDs []int) ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getClubsOrgs, clubIDs)
	if err != nil {
		return nil, wrapPostgresError(err)
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
			&c.IsAdmin,
			&c.ClubName,
			&c.ClubID,
		)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		oarr = append(oarr, c)
	}
	if len(oarr) == 0 {
		return nil, ErrPostgresNotFoundError
	}
	return oarr, nil
}

const getClubSubOrgs = `
SELECT
	orgs.role_name,
	orgs.role_spec,
	mem.id,
	mem.hash_password,
	mem.login,
	mem.media_id,
	mem.telegram,
	mem.vk,
	mem.name,
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
	WHERE id > 0
) as clubs
ON (club_org.club_id = clubs.id)
JOIN
(
	SELECT
        id,
		role_name,
		role_spec
	FROM club_role
    WHERE role_clearance = 2
) as orgs
ON orgs.id = club_org.role_id
WHERE club_id = ANY((SELECT id FROM club WHERE parent_id = $1)) AND club_id > 0
`

func (s *Postgres) GetClubSubOrgs(_ context.Context, clubID int) ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getClubSubOrgs, clubID)
	if err != nil {
		return nil, wrapPostgresError(err)
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
			&c.IsAdmin,
			&c.ClubName,
		)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		c.ClubID = clubID
		oarr = append(oarr, c)
	}
	if len(oarr) == 0 {
		return nil, ErrPostgresNotFoundError
	}
	return oarr, nil
}

const getAllClubOrgs = `
SELECT
	orgs.role_name,
	orgs.role_spec,
	member_id,
	club_id,
	mem.hash_password,
	mem.login,
	mem.media_id,
	mem.telegram,
	mem.vk,
	mem.name,
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
	WHERE id > 0
) as clubs
JOIN
(
	SELECT
        id,
		role_name,
		role_spec
	FROM club_role
    WHERE role_clearance = 2
) as orgs
ON orgs.id = club_org.role_id
ON (club_org.club_id = clubs.id)
`

func (s *Postgres) GetAllClubOrgs(_ context.Context) ([]domain.ClubOrg, error) {
	oarr := []domain.ClubOrg{}
	rows, err := s.db.Query(getAllClubOrgs)
	if err != nil {
		return nil, wrapPostgresError(err)
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
			&c.IsAdmin,
			&c.ClubName,
		)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		oarr = append(oarr, c)
	}
	if len(oarr) == 0 {
		return nil, ErrPostgresNotFoundError
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
		return 0, wrapPostgresError(err)
	}
	return id, nil
}

const addOrgRole = `
INSERT INTO club_role (
    role_name,
    role_spec,
    role_clearance
) VALUES ($1, $2, 2)
RETURNING id
`

const addOrgs = `
INSERT INTO club_org (
	role_id,
    member_id,
    club_id,
) VALUES ($1, $2, $3)
`

func (s *Postgres) AddOrgs(_ context.Context, orgs []domain.ClubOrg) error {
	tx, err := s.db.Begin()
	if err != nil {
		return wrapPostgresError(err)
	}

	for _, org := range orgs {
		var id int
		err := s.db.QueryRow(addOrgRole, org.RoleName, org.RoleSpec).Scan(&id)
		if err != nil {
			tx.Rollback()
			return wrapPostgresError(err)
		}
		_, err = tx.Exec(addOrgs, id, org.ID, org.ClubID)
		if err != nil {
			tx.Rollback()
			return wrapPostgresError(err)
		}
	}

	return tx.Commit()
}

const getClubMediaFiles = `
SELECT
    id,
	ref_num,
	club_id,
	media_id
FROM club_photo
WHERE club_id = $1 AND club_id > 0
`

func (s *Postgres) GetClubMediaFiles(clubID int) ([]domain.ClubPhoto, error) {
	ph := []domain.ClubPhoto{}
	rows, err := s.db.Query(getClubMediaFiles, clubID)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		p := domain.ClubPhoto{}
		err := rows.Scan(&p.ID, &p.RefNumber, &p.ClubID, &p.MediaID)
		if err != nil {
			return nil, wrapPostgresError(err)
		}
		p.ClubID = clubID
		ph = append(ph, p)
	}
	if len(ph) == 0 {
		return nil, ErrPostgresNotFoundError
	}
	return ph, nil
}

const deleteClub = "DELETE FROM club WHERE id = $1"
const deleteClubOrgs = "DELETE FROM club_org WHERE club_id = $1"
const deleteClubMembers = "DELETE FROM member WHERE club_id = $1"
const deleteClubPhotos = "DELETE FROM club_photo WHERE club_id = $1"
const deleteClubEncounters = "DELETE FROM encounter WHERE club_id = $1"
const deleteClubDocuments = "DELETE FROM document WHERE club_id = $1"
const updateClubParents = "UPDATE club SET parent_id=null WHERE parent_id = $1"

func (s *Postgres) DeleteClubWithOrgs(_ context.Context, clubID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return wrapPostgresError(err)
	}

	_, err = tx.Exec(deleteClubOrgs, clubID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err)
	}

	_, err = tx.Exec(deleteClubPhotos, clubID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err)
	}

	_, err = tx.Exec(deleteClubEncounters, clubID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err)
	}

	_, err = tx.Exec(deleteClubDocuments, clubID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err)
	}

	_, err = tx.Exec(updateClubParents, clubID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err)
	}

	tag, err := tx.Exec(deleteClub, clubID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err)
	}

	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
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
WHERE id = $9 AND id > 0
`

func (s *Postgres) UpdateClub(_ context.Context, c *domain.Club, o []domain.ClubOrg) error {
	tx, err := s.db.Begin()
	if err != nil {
		return wrapPostgresError(err)
	}

	tag, err := tx.Exec(updateClub,
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
		return wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return ErrPostgresNotFoundError
	}

	_, err = tx.Exec(deleteClubOrgs, c.ID)
	if err != nil {
		tx.Rollback()
		return wrapPostgresError(err)
	}

	for _, org := range o {
		_, err = tx.Exec(addOrgs, org.RoleName, org.RoleSpec, org.ID, c.ID)
		if err != nil {
			tx.Rollback()
			return wrapPostgresError(err)
		}
	}
	err = tx.Commit()
	return wrapPostgresError(err)
}
