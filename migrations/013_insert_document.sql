-- +goose Up
-- +goose StatementBegin

INSERT INTO category (name)
VALUES
('Category 1'),
('Category 2'),
('Category 3');

INSERT INTO document (name, key, club_id, category_id)
VALUES
('1.pdf', '1/1.pdf', 1, 1),
('2.pdf', '2/2.pdf', 2, 2),
('3.pdf', '3/3.pdf', 3, 3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

truncate table document CASCADE;
truncate table category CASCADE;

-- +goose StatementEnd