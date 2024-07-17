-- +goose Up
-- +goose StatementBegin

INSERT INTO mediafile (id, name, image)
VALUES
(0, 'default', 'default');  -- Important, do not remove

insert into mediafile (name, key)
values
('1.jpg', '1.jpg'),
('2.jpg', '2.jpg'),
('3.jpg', '3.jpg'),
('4.jpg', '4.jpg'),
('5.jpg', '5.jpg'),
('6.jpg', '6.jpg'),
('7.jpg', '7.jpg'),
('8.jpg', '8.jpg'),
('9.jpg', '9.jpg'),
('10.jpg', '10.jpg'),
('11.jpg', '11.jpg');

insert into default_media (media_id)
VALUES (1), (2), (3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

truncate table mediafile CASCADE;

-- +goose StatementEnd