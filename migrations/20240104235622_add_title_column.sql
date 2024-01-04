-- +goose Up
-- +goose StatementBegin
ALTER TABLE events ADD COLUMN title text default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
