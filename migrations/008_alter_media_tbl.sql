-- +goose Up
-- +goose StatementBegin
ALTER table mediafile rename column image to image_url;

ALTER table mediafile ALTER column image_url type TEXT;
update mediafile set image_url = 'localhost:5000/about/arch.png';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
