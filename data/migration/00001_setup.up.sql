Create table if not exists "user" (
    id bigint primary key,
    email varchar not null,
    created_at timestamptz not null,
    updated_at timestamptz  not null,
    constraint uq_user_email unique (email)
)

Create table if not exists "friend" (
    id bigint primary key,
    user1_id bigint not null,
    user2_id bigint not null,
    created_at timestamptz not null,
    updated_at timestamptz  not null,
    constraint fk_friend_user_user1_id foreign key (user1_id) references user(id),
    constraint fk_friend_user_user2_id foreign key (user2_id) references user(id),
    constraint uq_friend unique (user1_id, user2_id)
)

Create table if not exists "subscribe" (
    id bigint primary key,
    requestor_id bigint not null,
    target_id bigint not null,
    created_at timestamptz not null,
    updated_at timestamptz  not null,
    constraint fk_subscribe_user_requestor_id foreign key (requestor_id) references user(id),
    constraint fk_subscribe_user_target_id foreign key (target_id) references user(id)
)

Create table if not exists "block" (
    id bigint primary key,
    requestor_id bigint not null,
    target_id bigint not null,
    created_at timestamptz not null,
    updated_at timestamptz  not null,
    constraint fk_block_user_requestor_id foreign key (requestor_id) references user(id),
    constraint fk_block_user_target_id foreign key (target_id) references user(id)
)
