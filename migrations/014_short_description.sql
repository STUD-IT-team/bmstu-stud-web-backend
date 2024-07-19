-- +goose Up
-- +goose StatementBegin

ALTER TABLE club ADD short_description text default '';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE club DROP short_description;

-- +goose StatementEnd