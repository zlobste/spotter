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
    salary   money not null default 0,
    primary key (id)
);

create table groups
(
    id          bigserial,
    title       text not null default '',
    description text not null default '',
    level       int  not null default 1,
    primary key (id)
);

create table timers
(
    id         bigserial,
    group_id   bigserial                not null,
    start_time timestamp with time zone not null,
    duration   timestamp                not null,
    primary key (id),
    foreign key (group_id) references groups (id) on delete cascade
);

create table confirmations
(
    user_id   bigserial                not null,
    timer_id  bigserial                not null,
    date      timestamp with time zone not null,
    confirmed bool                     not null default false,
    primary key (user_id, timer_id, date),
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (timer_id) references timers (id) on delete cascade
);

create table votings
(
    id          bigserial,
    type        int                      not null default 0,
    victim      bigserial                not null,
    title       text                     not null default '',
    description text                     not null default '',
    end_time    timestamp with time zone not null,

    primary key (id),
    foreign key (victim) references users (id) on delete cascade
);

create table votes
(
    user_id   bigserial not null,
    voting_id bigserial not null,
    decided   bool      not null default false,
    primary key (user_id, voting_id),
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (voting_id) references votings (id) on delete cascade
);

create table users_groups
(
    user_id  bigserial not null,
    group_id bigserial not null,
    primary key (user_id, group_id),
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (group_id) references groups (id) on delete cascade
);

-- +migrate Down
drop table users cascade;
drop table groups cascade;
drop table timers cascade;
drop table confirmations cascade;
drop table votings cascade;
drop table votes cascade;
drop table users_groups cascade;