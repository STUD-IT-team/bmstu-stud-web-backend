-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS member
(
    id       serial primary key,
    hash_password bytea not null,
    login    text  not null unique,
    media_id int not null,
    telegram text    default '',
    name     text    default '',
    role_id  int     default 0,
    is_admin boolean default false
);

create table IF NOT EXISTS member_role
(
    role_id   serial primary key,
    role_name text not null,
    role_spec text not null
);

CREATE table IF NOT EXISTS feed
(
    id         serial primary key,
    title      text      not null,
    approved   boolean default false,
    description text      not null,
    media_id  int    not null,
    vk_post_url text default '',
    updated_at timestamp not null,
    created_at timestamp not null,
    views      int     default 0,
    created_by int       not null
);

-- TODO: REDO in domain
create table if not exists mediafile (
    id serial primary key,
    name text default '',
    image base64 default ''
)

create table IF NOT EXISTS event
(
    id            serial primary key,
    title         text      not null,
    description   text      not null,
    prompt        text      not null,
    media_id      int       not null,
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
drop table IF EXISTS member_role;
drop table IF EXISTS  member;
drop table IF EXISTS feed;
drop table IF EXISTS event;
drop table IF EXISTS mediafile;
-- +goose StatementEnd
