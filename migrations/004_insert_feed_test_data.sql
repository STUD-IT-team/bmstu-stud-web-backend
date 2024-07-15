-- +goose Up
-- +goose StatementBegin

INSERT INTO member_role (role_name, role_spec) VALUES
('admin', 'admin'),
('user', 'user');

INSERT INTO member (hash_password, login, media_id, role_id)
VALUES
('1234', '1234', 1, 1),
('1', '1', 2, 2);

INSERT INTO feed (title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by)
VALUES
('11', true, '33', 1, '132', '2004-10-19 10:23:54', '2004-10-19 10:23:54', 13, 1),
('1', true, '33', 2, '132', '2004-10-19 10:23:54', '2004-10-19 10:23:54', 13, 1),
('22', false, '44', 3, '321', '2005-11-19 10:23:54', '1900-10-19 10:23:54', 14, 2);

INSERT INTO encounter (count, description, club_id)
VALUES
('1', 'kcoc', 0),
('11', 'kcoc', 0),
('2', 'sllab', 1),
('22', 'sllab', 2);

INSERT INTO event (title, description, prompt,  media_id,  date, approved, created_at, main_org, reg_url, reg_open_date, feedback_url, club_id)
VALUES
('kcoc', 'sllab', 'kcid', 1, '2005-11-19 10:23:54', true, '2005-11-19 10:23:54', 1, 'ahh', '2005-11-19 10:23:54', '123', 1),
('sinep', 'stun', 'nibor', 1, '2005-11-19 10:23:54', true, '2005-11-19 10:23:54', 2, 'ahh', '2005-11-19 10:23:54', '123', 0);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM encounter;
DELETE FROM feed;
DELETE FROM event;
DELETE FROM member;
DELETE FROM member_role;
-- +goose StatementEnd