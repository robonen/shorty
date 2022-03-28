create table links (
    id bigserial primary key,
    user_id bigint not null references users (id) on update cascade on delete no action,
    hash varchar(7) not null,
    original text not null,
    expiration_at timestamp default null,
    created_at timestamp not null,
    updated_at timestamp not null
);

create index hash_index on links (hash);