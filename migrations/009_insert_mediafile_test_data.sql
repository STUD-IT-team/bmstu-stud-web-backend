-- +goose Up
-- +goose StatementBegin

INSERT INTO mediafile (name, image)
VALUES
('cockballs.png', 'idef-0id9'),
('balls.jpg', 'idef-0912'),
('cock.pm4','ifas-axsx');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM mediafile;
-- +goose StatementEnd