-- +goose Up
-- +goose StatementBegin

INSERT INTO member (hash_password, login, media_id, telegram, vk, name, is_admin)
VALUES
('test', 'toha', 1, '@toha', 'vk.com/toha', 'Антон Павленко', true),
('$2a$10$kj0GwI3q1H0PgOzuLqK5uOhPPvA42upL8CdIm/4luikQBYNKVxXay', 'imp', 1, '@imp', 'vk.com/imp', 'Дмитрий Шахнович', true),
('test', 'dasha', 1, '@dasha', 'vk.com/dasha', 'Дарья Серышева', true),
('test', 'paioid', 1, '@paioid', 'vk.com/paioid', 'Андрей Поляков', true),
('test', 'admin', 1, '@admin', 'vk.com/admin', 'Админ', true),
('test', 'user', 1, '@user', 'vk.com/user', 'Юзер', true),
('test', 'user2', 1, '@user2', 'vk.com/user2', 'Юзер2', true);


INSERT INTO club_org (club_id, member_id, role_name, role_spec)
VALUES
(1, 1, 'Молодец', 'IT'),
(1, 1, 'Красава', 'NIT'),
(1, 2, 'Веселый', 'Finance'),
(1, 2, 'Хорошо', 'HR'),
(1, 3, 'Богатый', 'Marketing'),
(1, 3, 'Молодец', 'Sales'),
(2, 4, 'Умный', 'Engineering');

UPDATE club
SET parent_id = 1
WHERE id = 2;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM club_org;
DELETE FROM member;

-- +goose StatementEnd