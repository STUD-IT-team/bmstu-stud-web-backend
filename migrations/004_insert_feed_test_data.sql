-- +goose Up
-- +goose StatementBegin

INSERT INTO feed (title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by)
VALUES
('11', true, '33', 1, '132', '2004-10-19 10:23:54', '2004-10-19 10:23:54', 13, 11),
('1', true, '33', 1, '132', '2004-10-19 10:23:54', '2004-10-19 10:23:54', 13, 11),
('22', false, '44', 2, '321', '2005-11-19 10:23:54', '1900-10-19 10:23:54', 14, 12);

INSERT INTO encounter (count, description, club_id)
VALUES
('1', 'cock', 0),
('11', 'kcoc', 0),
('2', 'balls', 1),
('22', 'sllab', 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE * FROM feed;
-- +goose StatementEnd