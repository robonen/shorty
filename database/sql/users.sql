create table users (
    id bigserial primary key,
    name varchar(255) default null,
    email varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    logging_at timestamp not null
);

create index email_index on users (email);