-- +goose Up
-- +goose StatementBegin

ALTER TABLE mediafile RENAME column image_url TO key;

UPDATE mediafile set key = 'arch.png' where id=1;
UPDATE mediafile set key = '2.jpg' where id=2;
UPDATE mediafile set key = '3.jpg' where id=3;
UPDATE mediafile set key = '4.jpg' where id=4;
UPDATE mediafile set key = '5.jpg' where id=5;
UPDATE mediafile set key = '6.jpg' where id=6;
UPDATE mediafile set key = '7.jpg' where id=7;
UPDATE mediafile set key = '8.jpg' where id=8;
UPDATE mediafile set key = '9.jpg' where id=9;
INSERT INTO mediafile (key, name) VALUES
('10.jpg', 'image10.jpg'),
('11.jpg', 'image11.jpg'),
('12.jpg', 'image12.jpg');

ALTER TABLE mediafile ADD CONSTRAINT key_unique UNIQUE (key)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd