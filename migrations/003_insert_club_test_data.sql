-- +goose Up
-- +goose StatementBegin

INSERT INTO mediafile (id, name, image)
VALUES
(0, 'default', 'default');  -- Important, do not remove

INSERT INTO mediafile (name, image)
VALUES
('IT-dep.jpg', 'image1.jpg'),
('NIT-dep.jpg', 'image2.jpg'),
('Finance-dep.jpg', 'image3.jpg');

INSERT INTO club (id, logo) VALUES (0, 0);  -- Important, do not remove
INSERT INTO club (name, short_name, description, type, logo, vk_url, tg_url)
VALUES
('IT-department', 'IT-dep', 'Typo iT', 'IT', 1, 'vk.com', 'tg.me'),
('Not IT-department', 'NIT-dep', 'Typo Ne iT', 'IT', 1, 'vk.com', 'tg.me'),
('Finance-department', 'Finance-dep', 'Typo Fin', 'Finance', 2, 'vk.com', 'tg.me'),
('HR-department', 'HR-dep', 'Typo Hr', 'HR', 2, 'vk.com', 'tg.me'),
('Marketing-department', 'Marketing-dep', 'Typo Mark', 'Marketing' , 3, 'vk.com', 'tg.me'),
('Sales-department', 'Sales-dep', 'Typo Sa', 'Sales', 3, 'vk.com', 'tg.me'),
('Engineering-department', 'Engineering-dep', 'Typo Eng', 'Engineering', 3, 'vk.com', 'tg.me');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM club;
DELETE FROM mediafile;

-- +goose StatementEnd