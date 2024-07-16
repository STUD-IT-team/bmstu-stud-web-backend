-- +goose Up
-- +goose StatementBegin

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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table IF EXISTS club_org CASCADE;
drop table IF EXISTS club_photo CASCADE;
drop table IF EXISTS club_role CASCADE;
drop table IF EXISTS club CASCADE;
drop table IF EXISTS event CASCADE;
drop table IF EXISTS encounter CASCADE;

-- +goose StatementEnd