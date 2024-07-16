-- +goose Up
-- +goose StatementBegin

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

-- +goose StatementEnd
