INSERT INTO users (id, email, "password", created_at, updated_at) values 
(1, 'test1@example.com', 'test1', now(), now()),
(2, 'test2@example.com', 'test2', now(), now()),
(3, 'test3@example.com', 'test3', now(), now());

INSERT INTO relationships (id, requestor_id, target_id, "type", created_at, updated_at) values 
(1, 1, 2, 'FRIEND', now(), now()),
(2, 1, 3, 'FRIEND', now(), now());
