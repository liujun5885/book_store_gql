-- Deploy book_store_ugc:user_profile_and_settings to pg

BEGIN;

CREATE TABLE user_profiles
(
    user_id  UUID REFERENCES users ON DELETE CASCADE,
    address  VARCHAR(1024) NOT NULL DEFAULT '',
    city     VARCHAR(256)  NOT NULL DEFAULT '',
    province VARCHAR(256)  NOT NULL DEFAULT '',
    country  VARCHAR(256)  NOT NULL DEFAULT '',
    job      VARCHAR(256)  NOT NULL DEFAULT '',
    school   VARCHAR(256)  NOT NULL DEFAULT ''
);

CREATE TABLE user_settings
(
    user_id        UUID REFERENCES users ON DELETE CASCADE,
    kindle_account VARCHAR(256) NOT NULL DEFAULT ''
);

COMMIT;
