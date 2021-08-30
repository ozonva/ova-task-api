-- +goose Up
CREATE TYPE task_status AS ENUM ('new', 'active', 'resolved');
CREATE table tasks
(
    id          serial primary key,
    title       text,
    description text,
    status      task_status,
    created_at  timestamptz
);
-- +goose Down
