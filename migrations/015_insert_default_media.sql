-- +goose Up
-- +goose StatementBegin
INSERT INTO default_media (media_id)
VALUES 
(1),
(1),
(1),
(1),
(1),
(1),
(1),
(1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM default_media;
-- +goose StatementEnd
