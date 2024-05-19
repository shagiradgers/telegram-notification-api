-- +goose Up
create table if not exists users
(
    id                  bigint generated always as identity primary key not null,
    telegram_id         bigint                                          not null,
    role                text                                            not null,
    notification_status text                                            not null,
    user_group          text                                            not null,
    firstname           text                                            not null,
    surname             text                                            not null,
    patronymic          text,
    mobile_phone        text                                            not null,
    status              text                                            not null
);
