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
