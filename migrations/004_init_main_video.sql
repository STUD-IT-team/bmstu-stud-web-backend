-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS main_video
(
    id      serial PRIMARY KEY,
    name    text DEFAULT '',
    key     text DEFAULT '',
    club_id int  NOT NULL,
    current boolean default NULL,
    FOREIGN KEY (club_id) REFERENCES club(id),
    CONSTRAINT true_1_per_club UNIQUE (club_id, current)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table IF EXISTS main_video CASCADE;

-- +goose StatementEnd