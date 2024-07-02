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
