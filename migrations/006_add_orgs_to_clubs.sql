-- +goose Up
-- +goose StatementBegin

INSERT INTO member (hash_password, login, media_id, telegram, vk, name, role_id, is_admin)
VALUES
('test', 'toha', 1, '@toha', 'vk.com/toha', 'Антон Павленко', 1, true),
('$2a$10$kj0GwI3q1H0PgOzuLqK5uOhPPvA42upL8CdIm/4luikQBYNKVxXay', 'imp', 1, '@imp', 'vk.com/imp', 'Дмитрий Шахнович', 1, true),
('test', 'dasha', 1, '@dasha', 'vk.com/dasha', 'Дарья Серышева', 1, true),
('test', 'paioid', 1, '@paioid', 'vk.com/paioid', 'Андрей Поляков', 1, true),
('test', 'admin', 1, '@admin', 'vk.com/admin', 'Админ', 1, true),
('test', 'user', 1, '@user', 'vk.com/user', 'Юзер', 1, true),
('test', 'user2', 1, '@user2', 'vk.com/user2', 'Юзер2', 1, true);


INSERT INTO club_org (club_id, member_id, role_name, role_spec)
VALUES
(1, 1, 'Молодец', 'IT'),
(2, 2, 'Красава', 'NIT'),
(3, 3, 'Веселый', 'Finance'),
(4, 4, 'Хорошо', 'HR'),
(5, 5, 'Богатый', 'Marketing'),
(6, 6, 'Молодец', 'Sales'),
(7, 7, 'Умный', 'Engineering');

UPDATE club
SET parent_id = 1
WHERE id = 2;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM member;
DELETE FROM club_org;

-- +goose StatementEnd