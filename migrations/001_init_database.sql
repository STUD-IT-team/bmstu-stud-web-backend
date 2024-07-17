-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS member
(
    id       serial primary key,
    hash_password bytea not null,
    login    text    not null unique,
    media_id int     not null,
    telegram text    default '',
    vk       text    default '',
    name     text    default '',
    is_admin boolean default false
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
    key text not null unique
);

alter table member add foreign key (media_id) references mediafile(id);

alter table feed add foreign key (media_id) references mediafile(id);
alter table feed add foreign key (created_by) references member(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table IF EXISTS mediafile CASCADE;
drop table IF EXISTS feed CASCADE;
drop table IF EXISTS member CASCADE;

-- +goose StatementEnd