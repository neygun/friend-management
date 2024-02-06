INSERT INTO users (id, email, "password", created_at, updated_at) values 
(1, 'test1@example.com', 'test1', now(), now()),
(2, 'test2@example.com', 'test2', now(), now()),
(3, 'test3@example.com', 'test3', now(), now()),
(4, 'test4@example.com', 'test4', now(), now()),
(5, 'test5@example.com', 'test5', now(), now());

INSERT INTO relationships (id, requestor_id, target_id, "type", created_at, updated_at) values 
(1, 1, 4, 'FRIEND', now(), now()),
(2, 5, 1, 'SUBSCRIBE', now(), now()),
(3, 3, 1, 'BLOCK', now(), now());
