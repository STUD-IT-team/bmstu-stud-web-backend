-- +goose Up
-- +goose StatementBegin
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table IF EXISTS document CASCADE;
drop table IF EXISTS category CASCADE; 
-- +goose StatementEnd
