-- +goose Up
-- +goose StatementBegin

ALTER table mediafile rename column image_url to key;

update mediafile set key = 'idef-0id9';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM mediafile;
-- +goose StatementEnd