-- +goose Up
-- +goose StatementBegin

insert into feed (title, approved, description, media_id, vk_post_url, updated_at, created_at, views, created_by)
values
('РЕГИСТРАЦИЯ НА СОПК', true, '🔥 Сегодня стартовала приемная кампания 2024 года, и МГТУ им. Н.Э. Баумана радушно распахнул свои двери для будущих студентов. В мае прошло несколько этапов отбора сотрудников, и сегодня они уже помогают абитуриентам с подачей документов, выбором направлений и консультациями. Ректор — Михаил Валерьевич Гордин и и.о. проректора по молодежной работе и воспитательной деятельности — Дмитрий Андреевич Сулегин (https://t.me/sulegin_bmstu) обратились к сотрудникам с напутственными словами и зарядили ребят на продуктивную работу этим летом!🫶 Сотрудники СОПК хорошо понимают все страхи абитуриентов, ведь сами проходили через этап поступления. Ребята ждут вас в Университете, чтобы помочь и ответить на все вопросы. До скорой встречи в МГТУ им. Н.Э. Баумана!', 6, 'https://vk.com/wall-26724538_19086', current_date, current_date, 10, 1),
('РЕГИСТРАЦИЯ НА ШПШ', true, 'На ШМБ приезжают Бауманцы, которые превращают это мероприятие в незабываемое событие для всех первокурсников. Их работа начинается на «Школе Перед Школой», где они не просто знакомятся, но и становятся крепкой командой. Каждая роль здесь важна, и каждый участник вносит неоценимый вклад, выкладываясь на все 100%, чтобы ШМБ стало ярким и запоминающимся мероприятием для всех!',7, 'https://vk.com/wall-26724538_19086', current_date, current_date, 10, 1),
('РЕГИСТРАЦИЯ В СТРОЙ ОТРЯД!', true, '🔥 Сегодня стартовал набор в стройотряд ДрУжБа', 8, 'https://vk.com/wall-26724538_19086', current_date, current_date, 10, 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

truncate table feed CASCADE;

-- +goose StatementEnd