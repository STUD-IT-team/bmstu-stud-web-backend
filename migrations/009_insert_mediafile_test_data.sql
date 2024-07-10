-- +goose Up
-- +goose StatementBegin

ALTER TABLE mediafile RENAME column image_url TO key;

UPDATE mediafile set key = '1.jpg' where id=1;
UPDATE mediafile set key = '2.jpg' where id=2;
UPDATE mediafile set key = '3.jpg' where id=3;
UPDATE mediafile set key = '4.jpg' where id=4;
UPDATE mediafile set key = '5.jpg' where id=5;
UPDATE mediafile set key = '6.jpg' where id=6;
UPDATE mediafile set key = '7.jpg' where id=7;
UPDATE mediafile set key = '8.jpg' where id=8;
UPDATE mediafile set key = '9.jpg' where id=9;

ALTER TABLE mediafile ADD CONSTRAINT key_unique UNIQUE (key)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM mediafile;
-- +goose StatementEnd