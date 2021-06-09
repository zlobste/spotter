-- +migrate Up
create table users
(
    id       bigserial,
    name     text not null,
    surname  text not null,
    email    text not null unique,
    password text not null,
    role     int  not null default 2,
    blocked  bool not null default false,
    primary key (id)
);

create table timers
(
    id         bigserial,
    user_id    bigserial                not null,
    start_time timestamp with time zone not null,
    end_time   timestamp with time zone,
    pending    bool                     not null default false,
    primary key (id),
    foreign key (user_id) references users (id) on delete cascade
);

create table proofs
(
    id         bigserial,
    timer_id   bigserial                not null,
    time       timestamp with time zone not null,
    percentage float                    not null,
    confirmed  bool                     not null default false,
    primary key (id),
    foreign key (timer_id) references timers (id) on delete cascade
);

-- +migrate Down
drop table users cascade;
drop table timers cascade;
drop table proofs cascade;
