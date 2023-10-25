INSERT INTO users (id, email, created_at, updated_at) values (1, 'test1@example.com', now(), now());
INSERT INTO users (id, email, created_at, updated_at) values (2, 'test2@example.com', now(), now());
INSERT INTO users (id, email, created_at, updated_at) values (3, 'test3@example.com', now(), now());
INSERT INTO relationships (id, requestor_id, target_id, "type", created_at, updated_at) values (1, 1, 2, 'FRIEND', now(), now());
INSERT INTO relationships (id, requestor_id, target_id, "type", created_at, updated_at) values (2, 1, 3, 'FRIEND', now(), now());
