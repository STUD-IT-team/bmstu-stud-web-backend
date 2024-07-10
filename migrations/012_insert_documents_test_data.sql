-- +goose Up
-- +goose StatementBegin

INSERT INTO document (name, key, club_id)
VALUES
('admin', 'admin', 0),
('user', 'user', 1),
('manager','manager', 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM document;
-- +goose StatementEnd