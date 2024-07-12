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
    image text default ''
);

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
    reg_open_date timestamp not null, 
    feedback_url  text    default ''
);

alter table member add foreign key (media_id) references mediafile(id);
alter table member add foreign key (role_id) references member_role(role_id);

alter table feed add foreign key (media_id) references mediafile(id);
alter table feed add foreign key (created_by) references member(id);

alter table event add foreign key (media_id) references mediafile(id);
alter table event add foreign key (created_by) references member(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table IF EXISTS member_role CASCADE;
drop table IF EXISTS member CASCADE;
drop table IF EXISTS feed CASCADE;
drop table IF EXISTS event CASCADE;
drop table IF EXISTS mediafile CASCADE;
-- +goose StatementEnd
