-- +goose Up
-- +goose StatementBegin
CREATE table feed (
  id serial primary key,
  text_field text not null,
  media_url text default ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table feed;
-- +goose StatementEnd
