-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS members
(
    id       serial primary key,
    password bytea not null,
    login    text  not null unique,
    telegram text    default '',
    name     text    default '',
    role_id  int     default 0,
    is_admin boolean default false
);

create table IF NOT EXISTS member_roles
(
    role_id   serial primary key,
    role_name text not null,
    role_spec text not null
);

CREATE table IF NOT EXISTS posts
(
    id         serial primary key,
    approved   boolean default false,
    description text      not null,
    media_url  text    default '',
    updated_at timestamp not null,
    created_at timestamp not null,
    views      int     default 0,
    created_by int       not null
);

create table IF NOT EXISTS events
(
    id            serial primary key,
    description   text      not null,
    date          timestamp not null,
    approved      boolean default false,
    created_at    timestamp not null,
    created_by    int       not null,
    reg_url       text    default '',
    reg_open_date text    default '',
    feedback_url  text    default ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table IF EXISTS member_roles;
drop table IF EXISTS  members;
drop table IF EXISTS posts;
drop table IF EXISTS events;
-- +goose StatementEnd
