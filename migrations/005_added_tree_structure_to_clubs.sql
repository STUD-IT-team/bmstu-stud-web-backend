-- +goose Up
-- +goose StatementBegin

ALTER TABLE club ADD COLUMN parent_id not null;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE club DROP COLUMN parent_id;

-- +goose StatementEnd