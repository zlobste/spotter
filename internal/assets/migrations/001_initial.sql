-- +migrate Up
create table users
(
    id       bigserial,
    name     text  not null default '',
    surname  text  not null default '',
    email    text  not null unique,
    password text  not null,
    role     int   not null default 2,
    balance  money not null default 0,
    PRIMARY KEY (id)
);

-- +migrate Down
drop table users cascade;