-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS member_event
(
    id            serial primary key,
    event_id      int not null,
    member_id     int not null,
    role_id       int not null,
    division      text default ''
);

create table IF NOT EXISTS member_event_role
(
    id            serial primary key,
    name          text not null
);

alter table member_event add foreign key (event_id) references event(id);
alter table member_event add foreign key (member_id) references member(id);
alter table member_event add foreign key (role_id) references member_event_role(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table IF EXISTS member_event_role CASCADE;
drop table IF EXISTS member_event CASCADE;

-- +goose StatementEnd
