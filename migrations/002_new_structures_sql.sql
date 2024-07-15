-- +goose Up
-- +goose StatementBegin

create table if not exists encounter (
    id serial primary key,
    count text default '',
    description text default '',
    club_id int not null
);

create table if not exists club (
    id serial primary key,
    name text default '' unique,
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
    ref_num int default 0,
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

alter table encounter add foreign key (club_id) references club(id);

alter table club add FOREIGN KEY (logo) REFERENCES mediafile(id);
alter table event add foreign key (club_id) references club(id);

alter table club_photo add FOREIGN KEY (media_id) REFERENCES mediafile(id);
alter table club_photo add FOREIGN KEY (club_id) REFERENCES club(id);

alter table club_org add FOREIGN KEY (club_id) REFERENCES club(id);
alter table club_org add FOREIGN KEY (member_id) REFERENCES member(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists encounter CASCADE;
drop table if exists club_photo CASCADE;
drop table if exists club_org CASCADE;
drop table if exists club CASCADE;
-- +goose StatementEnd
