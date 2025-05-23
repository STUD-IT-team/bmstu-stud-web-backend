-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS main_video
(
    id      serial PRIMARY KEY,
    name    text DEFAULT '',
    key     text DEFAULT '',
    club_id int  NOT NULL,
    current boolean default NULL,
    FOREIGN KEY (club_id) REFERENCES club(id)
);

CREATE UNIQUE INDEX ON main_video (club_id, current)
WHERE current = TRUE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table IF EXISTS main_video CASCADE;

-- +goose StatementEnd
