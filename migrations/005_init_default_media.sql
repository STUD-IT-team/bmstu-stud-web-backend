-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS default_media
(
    id      serial PRIMARY KEY,
    media_id int not null,
    FOREIGN KEY (media_id) REFERENCES mediafile(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table IF EXISTS default_media CASCADE;

-- +goose StatementEnd