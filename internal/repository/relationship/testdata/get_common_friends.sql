INSERT INTO users (id, email, created_at, updated_at) values 
(1, 'test1@example.com', now(), now()),
(2, 'test2@example.com', now(), now()),
(3, 'test3@example.com', now(), now()),
(4, 'test4@example.com', now(), now());

INSERT INTO relationships (id, requestor_id, target_id, "type", created_at, updated_at) values 
(1, 1, 3, 'FRIEND', now(), now()),
(2, 1, 4, 'FRIEND', now(), now()),
(3, 2, 3, 'FRIEND', now(), now()),
(4, 2, 4, 'FRIEND', now(), now());
