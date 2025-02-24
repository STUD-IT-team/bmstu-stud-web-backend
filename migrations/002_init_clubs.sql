-- +goose Up
-- +goose StatementBegin

create TABLE if not exists encounter (
    id serial primary KEY,
    count text default '',
    description text default '',
    club_id int not null
);

create TABLE IF NOT EXISTS event
(
    id            serial primary KEY,
    title         text      not null,
    description   text      not null,
    prompt        text      not null,
    media_id      int       not null,
    date          timestamp not null,
    approved      boolean default false,
    created_at    timestamp not null,
    created_by    int       not null,
    club_id       int       not null,
    main_org      int       not null,
    reg_url       text    default '',
    reg_open_date timestamp not null,
    feedback_url  text    default ''
);

create TABLE if not exists club (
    id serial primary KEY,
    name text default '' unique,
    short_name text default '',
    description text default '',
    type text default '',
    logo int not null,
    vk_url text default '',
    tg_url text default '',
    parent_id int default null
);

create TABLE IF NOT EXISTS club_role
(
    id   serial primary KEY,
    role_name text not null,
    role_spec text not null,
    role_clearance int default 0
);
-- 0 - гость
-- 1 - участник
-- 2 - руководитель
-- 3 - админ

-- TODO in domain
create TABLE if not exists club_photo (
    id serial primary KEY,
    ref_num int default 0,
    media_id int not null,
    club_id int not null
);

create TABLE if not exists club_org (
    id serial primary KEY,
    club_id int not null,
    member_id int not null,
    role_id int not null
);

ALTER TABLE encounter ADD FOREIGN KEY (club_id) REFERENCES club(id);

ALTER TABLE club ADD FOREIGN KEY (logo) REFERENCES mediafile(id);
ALTER TABLE club ADD FOREIGN KEY (parent_id) REFERENCES club(id);

ALTER TABLE event ADD FOREIGN KEY (club_id) REFERENCES club(id);
ALTER TABLE event ADD FOREIGN KEY (media_id) REFERENCES mediafile(id);
ALTER TABLE event ADD FOREIGN KEY (main_org) REFERENCES member(id);

ALTER TABLE club_photo ADD FOREIGN KEY (media_id) REFERENCES mediafile(id);
ALTER TABLE club_photo ADD FOREIGN KEY (club_id) REFERENCES club(id);
ALTER TABLE club_photo ADD CONSTRAINT media_club_ids_unique UNIQUE (club_id, media_id);

ALTER TABLE club_org ADD FOREIGN KEY (club_id) REFERENCES club(id);
ALTER TABLE club_org ADD FOREIGN KEY (member_id) REFERENCES member(id);
ALTER TABLE event ADD FOREIGN KEY (created_by) REFERENCES member(id);
ALTER TABLE club_org ADD FOREIGN KEY (role_id) REFERENCES club_role(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS club_role CASCADE;
DROP TABLE IF EXISTS club_org CASCADE;
DROP TABLE IF EXISTS club_photo CASCADE;
DROP TABLE IF EXISTS event CASCADE;
DROP TABLE IF EXISTS encounter CASCADE;
DROP TABLE IF EXISTS club CASCADE;

-- +goose StatementEnd
