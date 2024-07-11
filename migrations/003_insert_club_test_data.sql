-- +goose Up
-- +goose StatementBegin

INSERT INTO club (name, short_name, description, type, logo, vk_url, tg_url)
VALUES
('IT-department', 'IT-dep', 'Typo iT', 'IT', 1, 'vk.com', 'tg.me'),
('Not IT-department', 'NIT-dep', 'Typo Ne iT', 'IT', 1, 'vk.com', 'tg.me'),
('Finance-department', 'Finance-dep', 'Typo Fin', 'Finance', 1, 'vk.com', 'tg.me'),
('HR-department', 'HR-dep', 'Typo Hr', 'HR', 1, 'vk.com', 'tg.me'),
('Marketing-department', 'Marketing-dep', 'Typo Mark', 'Marketing' , 1, 'vk.com', 'tg.me'),
('Sales-department', 'Sales-dep', 'Typo Sa', 'Sales', 1, 'vk.com', 'tg.me'),
('Engineering-department', 'Engineering-dep', 'Typo Eng', 'Engineering', 1, 'vk.com', 'tg.me');

INSERT INTO mediafile (name, image)
VALUES
('IT-dep.jpg', 'image1.jpg'),
('NIT-dep.jpg', 'image2.jpg'),
('Finance-dep.jpg', 'image3.jpg');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM club;
DELETE FROM mediafile;
-- +goose StatementEnd