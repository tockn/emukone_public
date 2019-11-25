
-- +migrate Up
ALTER TABLE events DROP COLUMN end_time;
-- +migrate Down
