create  table if not exists reactions(
    id serial primary key,
    post_id int references posts(id) not null,
    reaction_giver_id int references users(id) not null,
    created_at timestamp default now(),
    unique(post_id,reaction_giver_id)
);