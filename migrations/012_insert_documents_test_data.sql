-- +goose Up
-- +goose StatementBegin

INSERT INTO category (name)
VALUES
('first'),
('second'),
('third');

INSERT INTO document (name, key, club_id, category_id)
VALUES
('admin', 'admin', 0, 1),
('user', 'user', 1, 1),
('manager','manager', 1, 2);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM document;
DELETE FROM category;
-- +goose StatementEnd