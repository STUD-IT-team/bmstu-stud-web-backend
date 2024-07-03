-- +goose Up
-- +goose StatementBegin

ALTER TABLE club ADD COLUMN parent_id int default null;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE club DROP COLUMN parent_id;

-- +goose StatementEnd