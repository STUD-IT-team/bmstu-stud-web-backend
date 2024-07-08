package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllEventQuery = `SELECT 
id           ,
title        ,
description  ,
prompt       ,
media_id     ,
date         ,
approved     ,
created_at   ,
created_by   ,
reg_url      ,
reg_open_date,
feedback_url
FROM event`

func (p *Postgres) GetAllEvents(_ context.Context) ([]domain.Event, error) {
	var events []domain.Event

	rows, err := p.db.Query(getAllEventQuery)
	if err != nil {
		return []domain.Event{}, err
	}

	for rows.Next() {
		var event domain.Event

		err = rows.Scan(
			&event.ID, &event.Title, &event.Description,
			&event.Propmt, &event.MediaID, &event.Date,
			&event.Approved, &event.CreatedAt, &event.CreatedBy,
			&event.RegUrl, &event.RegOpenDate, &event.FeedbackUrl)

		if err != nil {
			return []domain.Event{}, err
		}

		events = append(events, event)
	}

	if len(events) == 0 {
		return []domain.Event{}, fmt.Errorf("no events found")
	}

	return events, nil
}

const getEventQuery = `SELECT
id           ,
title        ,
description  ,
prompt       ,
media_id     ,
date         ,
approved     ,
created_at   ,
created_by   ,
reg_url      ,
reg_open_date,
feedback_url
FROM event WHERE id=$1`

func (p *Postgres) GetEvent(_ context.Context, id int) (domain.Event, error) {
	var event domain.Event

	err := p.db.QueryRow(getEventQuery, id).Scan(
		&event.ID, &event.Title, &event.Description,
		&event.Propmt, &event.MediaID, &event.Date,
		&event.Approved, &event.CreatedAt, &event.CreatedBy,
		&event.RegUrl, &event.RegOpenDate, &event.FeedbackUrl)
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}

const postEventQuery = `INSERT INTO event (title, description, prompt,  media_id,  date, approved, created_at, created_by, reg_url, reg_open_date, feedback_url)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

func (p *Postgres) PostEvent(_ context.Context, event domain.Event) error {
	_, err := p.db.Exec(postEventQuery,
		event.Title, event.Description,
		event.Propmt, event.MediaID, event.Date,
		event.Approved, event.CreatedAt, event.CreatedBy,
		event.RegUrl, event.RegOpenDate, event.FeedbackUrl,
	)

	if err != nil {
		return fmt.Errorf("can't insert event into postgres %w", err)
	}

	return nil
}

const deleteEventQuery = "DELETE FROM event WHERE id=$1"

func (p *Postgres) DeleteEvent(_ context.Context, id int) error {
	_, err := p.db.Exec(deleteEventQuery, id)
	if err != nil {
		return fmt.Errorf("can't delete event on postgres %w", err)
	}

	return nil
}

const updateEventQuery = `
UPDATE event SET  
title=$1,
description=$2,
prompt=$3,
media_id=$4,
date=$5,
approved=$6,
created_at=$7,
created_by=$8,
reg_url=$9,
reg_open_date=$10,
feedback_url=$11
WHERE id=$12`

func (p *Postgres) UpdateEvent(_ context.Context, event domain.Event) error {
	_, err := p.db.Exec(updateEventQuery,
		event.Title, event.Description,
		event.Propmt, event.MediaID, event.Date,
		event.Approved, event.CreatedAt, event.CreatedBy,
		event.RegUrl, event.RegOpenDate, event.FeedbackUrl,
		event.ID,
	)
	if err != nil {
		return fmt.Errorf("can't update event on postgres %w", err)
	}

	return nil
}
