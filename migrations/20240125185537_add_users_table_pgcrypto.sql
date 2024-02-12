-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS stud_users
(
    id            uuid not null primary key,
    email         text  unique default '',
    password      text  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table IF EXISTS stud_users;
-- +goose StatementEnd
