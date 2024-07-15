-- +goose Up
-- +goose StatementBegin

create table IF NOT EXISTS question_category
(
id serial primary key,
category text default ''
);

create table IF NOT EXISTS faq
(
    id serial primary key,
    question text default '',
    answer text default '',
    category_id int not null references question_category(id),
    club_id int not null
);





-- +goose StatementEnd
DELETE FROM faq;
DELETE FROM question_category;

-- +goose Down
-- +goose StatementBegin
drop table if exists faq;
drop table if exists question_category;

-- +goose StatementEnd