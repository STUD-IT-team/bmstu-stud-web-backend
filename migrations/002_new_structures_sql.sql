-- +goose Up
-- +goose StatementBegin

create table if not exists encounter (
    id serial primary key,
    count text default '',
    descriptions text default '',
    club_id int not null
);

create table if not exists club (
    id serial primary key,
    name text default '',
    short_name text default '',
    description text default '',
    type text default '',
    logo int not null,
    vk_url text default '',
    tg_url text default ''
);


-- TODO in domain
create table if not exists club_photo (
    id serial primary key,
    ref_num int unique default 0,
    media_id int not null,
    club_id int not null
);

create table if not exists club_org (
    id serial primary key,
    club_id int not null,
    member_id int not null,
    role_name text default '',
    role_spec text default ''
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists encounter;
drop table if exists club_photo;
drop table if exists club_org;
drop table if exists club;
-- +goose StatementEnd
