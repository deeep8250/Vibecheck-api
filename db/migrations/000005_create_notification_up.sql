create table if not exists notifications(
    id serial primary key,
    user_id int references users(id) not null,
    reacted_user_id int references users(id) not null,
    reacted_post_id int references posts(id) not null,
    
    created_at timestamp default now(),
    unique(user_id,reacted_post_id)
); 