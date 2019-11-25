
-- +migrate Up
INSERT INTO events(id, name, start_time, end_time, why_description, free_description, place, ticket_price, header_image_url, meta_description, open_time) VALUES (1, 'event1', '2019-02-22 19:30', '2019-02-22 22:00', 'why', 'free', 'shimokita', '4000', 'https://pbs.twimg.com/media/D0VHDYYU8AExZbF.jpg', 'すげぇイベント', '2019-02-22 19:00');
INSERT INTO events_related_users(user_id, event_id) VALUES (3, 1);

-- +migrate Down
