-- Verify book_store_ugc:user_profile_and_settings on pg

BEGIN;

SELECT * FROM user_profiles;
SELECT * FROM user_settings;

ROLLBACK;
