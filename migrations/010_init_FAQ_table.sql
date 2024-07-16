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

INSERT INTO question_category(id, category)
VALUES
(1,'частые и полезные вопросы'),
(2,'бесполезные вопросы');

INSERT INTO faq(question, answer, category_id, club_id)
VALUES
('как какать?', 'нинаю', 1, 2),
('как писать код?', 'никто тебе не скажет', 1, 3),
('как отчислиться?', 'нильзя', 1, 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM faq;
DELETE FROM question_category;
-- +goose StatementEnd