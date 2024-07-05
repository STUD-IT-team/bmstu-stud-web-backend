-- +goose Up
-- +goose StatementBegin

ALTER table mediafile ALTER column image type TEXT;
update mediafile set image = 'localhost:5000/about/arch.png';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
