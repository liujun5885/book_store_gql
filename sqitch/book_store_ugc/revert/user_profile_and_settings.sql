-- Revert book_store_ugc:user_profile_and_settings from pg

BEGIN;

DROP TABLE user_profiles;
DROP TABLE user_settings;

COMMIT;
