BEGIN;
INSERT INTO test.customer
(id, created_at, updated_at, deleted_at, "name", phone, email, street, city, state, 
profile_pic)
VALUES('0189fabc-1afc-49f9-bf68-95453466b50d', now(), now(), 
NULL::timestamp with time zone, 'max', '123', 
'max@email.com', 'street max', 'city max', 'state max', 'profile max');

INSERT INTO test.customer
(id, created_at, updated_at, deleted_at, "name", phone, email, street, city, state, 
profile_pic)
VALUES('b5370f68-8e89-4ad6-97b2-893b0e9121ea', now(), now(), 
NULL::timestamp with time zone, 'fred', '234', 
'fred@email.com', 'street fred', 'city fred', 'state fred', 'profile fred');

INSERT INTO test.customer
(id, created_at, updated_at, deleted_at, "name", phone, email, street, city, state, 
profile_pic)
VALUES('a62d4b19-6161-4c05-ae77-dd0a5a10a9b6', now(), now(), 
NULL::timestamp with time zone, 'pam', '908', 
'pam@email.com', 'street pam', 'city pam', 'state pam', 'profile pam');


INSERT INTO test.category
(id, created_at, updated_at, deleted_at, "name", parent_category)
VALUES('b81c66b2-7d7a-4d80-a091-ee82dc796e67', now(), now(), NULL::timestamp with time zone, 'FMCG', NULL);

INSERT INTO test.category
(id, created_at, updated_at, deleted_at, "name", parent_category)
VALUES('19b28005-aa07-43f0-84d4-f722f48bfa29', now(), now(), NULL::timestamp with time zone, 'Soft drink', 'b81c66b2-7d7a-4d80-a091-ee82dc796e67');

INSERT INTO test.category
(id, created_at, updated_at, deleted_at, "name", parent_category)
VALUES('3298d05a-e4d5-46c8-958a-9e7fecd3fe48', now(), now(), NULL::timestamp with time zone, 'carbonated drink', '19b28005-aa07-43f0-84d4-f722f48bfa29');

INSERT INTO test.product
(id, created_at, updated_at, deleted_at, "name", 
image, description, unit_price, category_id)
VALUES('d0ce7618-0f51-4a33-96a4-895115886f39', now(), now(), NULL::timestamp with time zone, 
'pepsi', 'p1 image', 'p1111', 60, '3298d05a-e4d5-46c8-958a-9e7fecd3fe48');

INSERT INTO test.product
(id, created_at, updated_at, deleted_at, "name", 
image, description, unit_price, category_id)
VALUES('a4bc75b1-e585-4ac8-941f-f15a725f959e', now(), now(), NULL::timestamp with time zone, 
'maggi', 'maggi image', 'maggiiiii', 25, 'b81c66b2-7d7a-4d80-a091-ee82dc796e67');


INSERT INTO test."order"
(id, created_at, updated_at, deleted_at, customer_id, total_quantity,total_price)
VALUES('951c54e9-4b64-42fe-9d56-e8a9babc3f89', now(), now(), 
NULL::timestamp with time zone, '0189fabc-1afc-49f9-bf68-95453466b50d', 15,725);

INSERT INTO test."order"
(id, created_at, updated_at, deleted_at, customer_id, total_quantity,total_price)
VALUES('738a9ccd-2779-472f-8c1a-7d830ad3867f', now(), now(), 
NULL::timestamp with time zone, 'b5370f68-8e89-4ad6-97b2-893b0e9121ea', 5,300);

INSERT INTO test.order_item
(id, created_at, updated_at, deleted_at, order_id, product_id, 
quantity, price, row_total)
VALUES('6572120b-08cf-42bf-bae4-5970edc4a12b', now(), now(), 
NULL::timestamp with time zone, '951c54e9-4b64-42fe-9d56-e8a9babc3f89', 
'd0ce7618-0f51-4a33-96a4-895115886f39', 10, 60, 600);

INSERT INTO test.order_item
(id, created_at, updated_at, deleted_at, order_id, product_id, quantity, price, row_total)
VALUES('f04b6390-6662-4434-a4f0-d853754c8541', now(), now(), NULL::timestamp with time zone, 
'951c54e9-4b64-42fe-9d56-e8a9babc3f89', 'a4bc75b1-e585-4ac8-941f-f15a725f959e', 5, 25, 125);

INSERT INTO test.order_item
(id, created_at, updated_at, deleted_at, order_id, product_id, quantity, price, row_total)
VALUES('69c84285-8c7f-4e70-ac65-a218ef35c606', now(), now(), NULL::timestamp with time zone, 
'738a9ccd-2779-472f-8c1a-7d830ad3867f', 'd0ce7618-0f51-4a33-96a4-895115886f39', 5, 60, 300);

COMMIT;