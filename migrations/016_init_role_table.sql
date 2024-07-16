-- +goose Up
-- +goose StatementBegin


ALTER TABLE member_role RENAME TO club_role;
ALTER TABLE member DROP COLUMN role_id;
ALTER TABLE club_org DROP COLUMN role_spec;
ALTER TABLE club_org DROP COLUMN role_name;
ALTER TABLE club_role RENAME COLUMN role_id TO id;
ALTER TABLE club_org ADD COLUMN role_id int references club_role(id);

-- 0 - гость
-- 1 - участник
-- 2 - руководитель
-- 3 - админ
ALTER TABLE club_role ADD COLUMN role_clearance int;


INSERT INTO club_role (role_name, role_spec, role_clearance)
VALUES
('IT бог', 'Глава ИТЭ', 2),
('Программист', 'Программист', 2),
('Тестировщик', 'Тестировщик', 2),
('Бухгалтер', 'Бухгалтер', 2),
('Старший участник', 'Старший участник', 2),
('Младший участник', 'Младший участник', 2),
('Гость', 'Гость', 2);

UPDATE club_org SET role_id = 1 WHERE id = 1;
UPDATE club_org SET role_id = 2 WHERE id = 2;
UPDATE club_org SET role_id = 3 WHERE id = 3;
UPDATE club_org SET role_id = 4 WHERE id = 4;
UPDATE club_org SET role_id = 5 WHERE id = 5;
UPDATE club_org SET role_id = 6 WHERE id = 6;
UPDATE club_org SET role_id = 7 WHERE id = 7;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE member Add COLUMN role_id int;
-- +goose StatementEnd
