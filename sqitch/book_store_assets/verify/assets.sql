-- Verify book_store_assets:assets on pg

BEGIN;

SELECT * FROM books limit 1;
SELECT * FROM authors limit 1;
SELECT * FROM publishers limit 1;
SELECT * FROM book_authors limit 1;
SELECT * FROM book_publishers limit 1;

ROLLBACK;
