-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS users
(
    id            uuid not null primary key,
    email         text  default '',
    hash_password text  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table IF EXISTS users;
-- +goose StatementEnd
