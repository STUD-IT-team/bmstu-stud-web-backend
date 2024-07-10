-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS document
(
    id      serial primary key,
    name    text default '',
    key     text default '',
    club_id int  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table IF EXISTS document;
-- +goose StatementEnd
