Create table if not exists "user" (
    "id" bigint primary key,
    "email" varchar not null,
    "created_at" timestamptz not null,
    "updated_at" timestamptz  not null,
    constraint uq_user_email unique ("email")
);

Create table if not exists "relationship" (
    "id" bigint primary key,
    "requestor_id" bigint not null,
    "target_id" bigint not null,
    "type" varchar not null,
    "created_at" timestamptz not null,
    "updated_at" timestamptz  not null,
    constraint fk_subscribe_user_requestor_id foreign key ("requestor_id") references "user"("id"),
    constraint fk_subscribe_user_target_id foreign key ("target_id") references "user"("id"),
    constraint uq_relationship unique ("requestor_id", "target_id", "type")
);
