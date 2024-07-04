-- +goose Up
-- +goose StatementBegin

INSERT INTO club_photo (club_id, media_id, ref_num)
VALUES
(1, 1, 1),
(1, 2, 2),
(2, 3, 1),
(2, 4, 2),
(3, 5, 1),
(3, 6, 2),
(3, 7, 3),
(3, 8, 4),
(3, 9, 5);


INSERT INTO mediafile (name, image)
VALUES
('4.jpg', 'image4.jpg'),
('5.jpg', 'image5.jpg'),
('6.jpg', 'image6.jpg'),
('7.jpg', 'image7.jpg'),
('8.jpg', 'image8.jpg'),
('9.jpg', 'image9.jpg');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM club_photo;

-- +goose StatementEnd