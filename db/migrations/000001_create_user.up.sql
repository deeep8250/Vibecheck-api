create table if not exists users(
    id serial  primary key,
    username varchar(200) not null,
    email varchar(200) unique  not null,
    password_hash varchar(200) not null,
    created_at timestamp default now()

);