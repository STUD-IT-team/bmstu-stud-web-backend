-- +goose Up
-- +goose StatementBegin

INSERT INTO question_category(category)
VALUES
('частые и полезные вопросы'),
('бесполезные вопросы');

INSERT INTO faq(question, answer, category_id, club_id)
VALUES
('как какать?', 'нинаю', 1, 2),
('как писать код?', 'никто тебе не скажет', 1, 3),
('как отчислиться?', 'нильзя', 2, 1);

-- +goose StatementEnd
DELETE FROM faq;
DELETE FROM question_category;

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd