-- Revert book_store_assets:assets from pg

BEGIN;

DROP TABLE book_topics;
DROP TABLE book_authors;
DROP TABLE book_publishers;
DROP TABLE books;
DROP TABLE authors;
DROP TABLE publishers;
DROP TABLE topics;
DROP EXTENSION IF EXISTS "pg_trgm";
DROP EXTENSION IF EXISTS "uuid-ossp";

COMMIT;
