create table if not exists posts(
    id serial primary key,
    user_id int references users(id) not null,
     content text not null,
    mood_tag varchar(200) not null,
    emoji varchar(10) not null,
    post_date date default current_date,
    created_at timestamp default now(),
    unique(user_id,post_date)
);