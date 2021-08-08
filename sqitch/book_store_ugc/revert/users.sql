-- Revert book_store_ugc:users from pg

BEGIN;

DROP TABLE users;

COMMIT;
