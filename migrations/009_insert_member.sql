-- +goose Up
-- +goose StatementBegin

insert into member (hash_password, login, media_id, telegram, vk, name, is_admin)
values
('$2y$10$GvW6IPLIAda8K7cIVeQ02.Sh/s5hPi2QHXZRSxkwG0kZiG00qOyQi', 'TestHeadMaster', 1, '@testHeadTelegram', 'TestHeadTelegram', 'Антон Успенский', true),
('$2y$10$GvW6IPLIAda8K7cIVeQ02.Sh/s5hPi2QHXZRSxkwG0kZiG00qOyQi', 'TestTPmaster', 2, '@TestTPmaster', 'TestTPmasterTelegram', 'Максим Демьянов', false),
('$2y$10$GvW6IPLIAda8K7cIVeQ02.Sh/s5hPi2QHXZRSxkwG0kZiG00qOyQi', 'TestHeadMissis', 3, '@TestHeadMissis', 'TestHeadMissisTelegram', 'Екатерина Донскова', false),
('$2y$10$GvW6IPLIAda8K7cIVeQ02.Sh/s5hPi2QHXZRSxkwG0kZiG00qOyQi', '@testHeadGather', 4, '@testHeadGatherTelegram', 'testHeadGather', 'Ольга Вакулина', false),
('$2y$10$GvW6IPLIAda8K7cIVeQ02.Sh/s5hPi2QHXZRSxkwG0kZiG00qOyQi', 'TestHeadMedia', 5, '@testHeadMediaTelegram', 'testHeadMediaTelegram', 'Егор Федорук', false);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

truncate table member CASCADE;

-- +goose StatementEnd