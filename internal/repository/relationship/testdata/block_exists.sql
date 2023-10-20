TRUNCATE TABLE "user" CASCADE;
TRUNCATE TABLE "relationship";
INSERT INTO "user" (id, email, created_at, updated_at) values (1, 'test1@example.com', now(), now());
INSERT INTO "user" (id, email, created_at, updated_at) values (2, 'test2@example.com', now(), now());
INSERT INTO "relationship" (id, requestor_id, target_id, "type", created_at, updated_at) values (1, 1, 2, 'BLOCK', now(), now());
