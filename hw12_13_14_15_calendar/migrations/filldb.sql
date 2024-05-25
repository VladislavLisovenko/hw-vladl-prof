create table if not exists events
(
    id varchar(36) primary key,
    title varchar(150) not null default '',
    event_date date NOT NULL DEFAULT now(),
    expiration_date date NOT NULL DEFAULT now(),
    description text not null default '',
    user_id varchar(36) not null,
    seconds_until_notification integer default 0
);