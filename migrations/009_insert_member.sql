-- +goose Up
-- +goose StatementBegin

insert into member (hash_password, login, media_id, telegram, vk, name, is_admin)
values
('12345678', 'TestHeadMaster', 1, '@testHeadTelegram', 'TestHeadTelegram', 'Антон Успенский', true),
('12345678', 'TestTPmaster', 2, '@TestTPmaster', 'TestTPmasterTelegram', 'Максим Демьянов', false),
('12345678', 'TestHeadMissis', 3, '@TestHeadMissis', 'TestHeadMissisTelegram', 'Екатерина Донскова', false),
('12345678', '@testHeadGather', 4, '@testHeadGatherTelegram', 'testHeadGather', 'Ольга Вакулина', false),
('12345678', 'TestHeadMedia', 5, '@testHeadMediaTelegram', 'testHeadMediaTelegram', 'Егор Федорук', false);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

truncate table member CASCADE;

-- +goose StatementEnd