create table if not exists follows(
    id serial primary key,
    follower_id int references users(id) not null,
    followed_user_id  int references users(id) not null,
    created_at timestamp default now(),
    unique(follower_id,followed_user_id)    
    );