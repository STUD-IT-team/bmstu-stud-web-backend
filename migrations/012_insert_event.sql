-- +goose Up
-- +goose StatementBegin

-- INSERT INTO event (title, description, prompt, media_id, date, approved, created_at, club_id, main_org, reg_url, reg_open_date, feedback_url)
-- VALUES
-- ('Event 1', 'Description of Event 1', 'Prompt for Event 1', 1, '2022-01-01 10:00:00', true, '2022-01-01 09:00:00', 1, 1, 'https://example.com/register1', '2022-01-01 08:00:00', 'https://example.com/feedback1'),
-- ('Event 2', 'Description of Event 2', 'Prompt for Event 2', 2, '2022-02-01 10:00:00', false, '2022-02-01 09:00:00', 2, 2, 'https://example.com/register2', '2022-02-01 08:00:00', 'https://example.com/feedback2'),
-- ('Event 3', 'Description of Event 3', 'Prompt for Event 3', 3, '2022-03-01 10:00:00', true, '2022-03-01 09:00:00', 3, 3, 'https://example.com/register3', '2022-03-01 08:00:00', 'https://example.com/feedback3');
INSERT INTO event (title, description, prompt, media_id, date, approved, created_at, created_by, reg_url, reg_open_date, feedback_url)
VALUES
('Event 1', 'Description of Event 1', 'Prompt for Event 1', 1, '2022-01-01 10:00:00', true, '2022-01-01 09:00:00', 1, 'https://example.com/register1', '2022-01-01 08:00:00', 'https://example.com/feedback1'),
('Event 2', 'Description of Event 2', 'Prompt for Event 2', 2, '2022-02-01 10:00:00', false, '2022-02-01 09:00:00', 2, 'https://example.com/register2', '2022-02-01 08:00:00', 'https://example.com/feedback2'),
('Event 3', 'Description of Event 3', 'Prompt for Event 3', 3, '2022-03-01 10:00:00', true, '2022-03-01 09:00:00', 3, 'https://example.com/register3', '2022-03-01 08:00:00', 'https://example.com/feedback3');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

truncate table event CASCADE;

-- +goose StatementEnd