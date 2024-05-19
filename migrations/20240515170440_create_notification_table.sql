-- +goose Up
create table if not exists notifications
(
    id            bigint generated always as identity primary key not null,
    sender_id     bigint references users (id)                    not null,
    receiver_ids  bigint array                                    not null,
    message       text                                            not null,
    media_content text,
    status        text                                            not null,
    date          timestamp                                       not null
);

