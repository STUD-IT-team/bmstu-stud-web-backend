-- +goose Up
-- +goose StatementBegin

create table if not exists encounter (
    id serial primary key,
    count text default '',
    descriptions text default '',
    club_id int not null
);


-- TODO in domain
create table if not exists club_photo (
    id serial primary key,
    ref_num int unique default 0,
    media_id int references media(id),
    club_id int not null
)

create table if not exists club_org (
    id serial primary key,
    club_id int not null,
    member_id int not null,
    role_name text default '',
    role_spec text default
)


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists encounter
drop table if exists club_photo
drop table if exists club_org
-- +goose StatementEnd
