-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table IF NOT EXISTS users
(
    id            uuid not null default uuid_generate_v4() primary key,
    email         text  default '',
    hash_password text  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop EXTENSION IF EXISTS pgcrypto;
drop EXTENSION IF EXISTS "uuid-ossp";
drop table IF EXISTS users;
-- +goose StatementEnd
