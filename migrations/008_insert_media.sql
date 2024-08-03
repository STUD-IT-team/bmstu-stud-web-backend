-- +goose Up
-- +goose StatementBegin

INSERT INTO mediafile (id, name, key)
VALUES
(0, '', '');  -- Important, do not remove

insert into mediafile (name, key)
values
('1.jpg', '1.jpg'),  -- 1
('2.jpg', '2.jpg'),  -- 2
('3.jpg', '3.jpg'),  -- 3
('4.jpg', '4.jpg'),  -- 4
('5.jpg', '5.jpg'),  -- 5
('6.jpg', '6.jpg'),  -- 6
('7.jpg', '7.jpg'),  -- 7
('8.jpg', '8.jpg'),  -- 8
('9.jpg', '9.jpg'),  -- 9
('10.jpg', '10.jpg'),  -- 10
('11.jpg', '11.jpg');  -- 11

-- events
insert into mediafile (name, key)
values
('SHMB.png', 'events/SHMB.png'),  -- 12
('posv.jpg', 'events/posv.jpg'),  -- 13
('star.jpg', 'events/star.jpg'),  -- 14
('doors.jpg', 'events/doors.jpg'),  -- 15
('meloch.jpg', 'events/meloch.jpg'),  -- 16
('leg.jpg', 'events/leg.jpg'),  -- 17
('SSO.jpg', 'events/SSO.jpg'),  -- 18
('hard.jpg', 'events/hard.jpg'),  -- 19
('kino.jpg', 'events/kino.jpg'),  -- 20
('buissn.jpg', 'events/buissn.jpg'),  -- 21
('new_year.jpg', 'events/new_year.jpg'),  -- 22
('new_year_go_away.jpg', 'events/new_year_go_away.jpg'),  -- 23
('secret.jpg', 'events/secret.jpg'),  -- 24
('bars.jpg', 'events/bars.jpg'),  -- 25
('media.jpg', 'events/media.jpg'),  -- 26
('it_quest.jpg', 'events/it_quest.jpg'),  -- 27
('butter.jpg', 'events/butter.jpg'),  -- 28
('kwizon.jpg', 'events/kwizon.jpg'),  -- 29
('leader.jpg', 'events/leader.jpg'),  -- 30
('it_swag.jpg', 'events/it_swag.jpg'),  -- 31
('studakiada.jpg', 'events/studakiada.jpg'),  -- 32
('miska.jpg', 'events/miska.jpg'),  -- 33
('sopka.jpg', 'events/sopka.jpg'),  -- 34
('coord.jpg', 'events/coord.jpg'),  -- 35
('stud_vip.jpg', 'events/stud_vip.jpg');  -- 36


insert into default_media (media_id)
VALUES (1), (2), (3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

truncate table mediafile CASCADE;

-- +goose StatementEnd