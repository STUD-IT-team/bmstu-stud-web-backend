-- +goose Up
-- +goose StatementBegin

ALTER TABLE club ADD COLUMN parent_id int default null;
alter table club add FOREIGN KEY (parent_id) REFERENCES club(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE club DROP COLUMN parent_id;

-- +goose StatementEnd