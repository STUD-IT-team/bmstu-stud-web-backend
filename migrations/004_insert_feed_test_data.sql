-- +goose Up
-- +goose StatementBegin

INSERT INTO feed (title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by)
VALUES
('11', true, '33', 1, '132', '2004-10-19 10:23:54', '2004-10-19 10:23:54', 13, 11),
('1', true, '33', 1, '132', '2004-10-19 10:23:54', '2004-10-19 10:23:54', 13, 11),
('22', false, '44', 1, '321', '2005-11-19 10:23:54', '1900-10-19 10:23:54', 14, 12);

INSERT INTO encounter (count, description, club_id)
VALUES
('1', 'kcoc', 0),
('11', 'kcoc', 0),
('2', 'sllab', 1),
('22', 'sllab', 1);

INSERT INTO event (title, description, prompt,  media_id,  date, approved, created_at, created_by, reg_url, reg_open_date, feedback_url)
VALUES
('kcoc', 'sllab', 'kcid', 1, '2005-11-19 10:23:54', true, '2005-11-19 10:23:54', 3, 'ahh', '2005-11-19 10:23:54', '123'),
('sinep', 'stun', 'nibor', 1, '2005-11-19 10:23:54', true, '2005-11-19 10:23:54', 3, 'ahh', '2005-11-19 10:23:54', '123');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM feed;
DELETE FROM encounter;
DELETE FROM event;
-- +goose StatementEnd