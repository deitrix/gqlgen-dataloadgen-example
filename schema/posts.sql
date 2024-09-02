create table posts
(
    id      int auto_increment primary key,
    title   varchar(64) not null,
    text    text        not null,
    user_id int         null,
    constraint posts_users_id_fk
        foreign key (user_id) references users (id)
            on update set null on delete set null
);

