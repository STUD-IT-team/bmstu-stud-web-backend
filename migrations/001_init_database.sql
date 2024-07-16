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

-----------------------------------------------------

create table if not exists encounter (
    id serial primary key,
    count text default '',
    description text default '',
    club_id int not null
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
    club_id       int       not null,
    main_org      int       not null,
    reg_url       text    default '',
    reg_open_date timestamp not null, 
    feedback_url  text    default ''
);

create table if not exists club (
    id serial primary key,
    name text default '' unique,
    short_name text default '',
    description text default '',
    type text default '',
    logo int not null,
    vk_url text default '',
    tg_url text default '',
    parent_id int default null
);

create table IF NOT EXISTS club_role
(
    id   serial primary key,
    role_name text not null,
    role_spec text not null,
    role_clearance int default 0
);
-- 0 - гость
-- 1 - участник
-- 2 - руководитель
-- 3 - админ

-- TODO in domain
create table if not exists club_photo (
    id serial primary key,
    ref_num int default 0,
    media_id int not null,
    club_id int not null
);

create table if not exists club_org (
    id serial primary key,
    club_id int not null,
    member_id int not null,
    role_id int not null
);

alter table encounter add foreign key (club_id) references club(id);

alter table club add FOREIGN KEY (logo) REFERENCES mediafile(id);
alter table club add FOREIGN KEY (parent_id) REFERENCES club(id);

alter table event add foreign key (club_id) references club(id);
alter table event add foreign key (media_id) references mediafile(id);
alter table event add foreign key (main_org) references member(id);

alter table club_photo add FOREIGN KEY (media_id) REFERENCES mediafile(id);
alter table club_photo add FOREIGN KEY (club_id) REFERENCES club(id);

alter table club_org add FOREIGN KEY (club_id) REFERENCES club(id);
alter table club_org add FOREIGN KEY (member_id) REFERENCES member(id);
alter table club_org add FOREIGN KEY (role_id) REFERENCES club_role(id);

-----------------------------------------------------

create table IF NOT EXISTS document
(
    id      serial primary key,
    name    text default '',
    key     text default '' unique,
    club_id int  not null,
    category_id int not null
);

create table IF NOT EXISTS category
(
    id   serial primary key,
    name text default '' unique
);

alter table document add foreign key (category_id) references category(id);
alter table document add foreign key (club_id) references club(id);

-----------------------------------------------------

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

-----------------------------------------------------

create table IF NOT EXISTS default_media
(
    id      serial PRIMARY KEY,
    media_id int not null,
    FOREIGN KEY (media_id) REFERENCES mediafile(id)
);

-----------------------------------------------------

create table IF NOT EXISTS event_member
(
    id            serial primary key,
    event_id      int not null,
    member_id     int not null,
    role_id       int not null,
    division      text default ''
);

create table IF NOT EXISTS event_member_role
(
    id            serial primary key,
    name          text not null
);

alter table event_member add foreign key (event_id) references event(id);
alter table event_member add foreign key (member_id) references member(id);
alter table event_member add foreign key (role_id) references event_member_role(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table IF EXISTS event_member_role CASCADE;
drop table IF EXISTS event_member CASCADE;

----------------------

drop table IF EXISTS default_media CASCADE;

----------------------

drop table IF EXISTS main_video CASCADE;

----------------------

drop table IF EXISTS category CASCADE;
drop table IF EXISTS document CASCADE;

----------------------

drop table IF EXISTS club_org CASCADE;
drop table IF EXISTS club_photo CASCADE;
drop table IF EXISTS club_role CASCADE;
drop table IF EXISTS club CASCADE;
drop table IF EXISTS event CASCADE;
drop table IF EXISTS encounter CASCADE;


-----------------------

drop table IF EXISTS mediafile CASCADE;
drop table IF EXISTS feed CASCADE;
drop table IF EXISTS member CASCADE;

-- +goose StatementEnd