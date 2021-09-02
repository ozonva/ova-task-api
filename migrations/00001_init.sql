-- +goose Up
CREATE TABLE tasks
(
    id          serial primary key,
    userid      integer     not null,
    description text        not null,
    created_at  timestamptz not null
);
-- +goose Down
