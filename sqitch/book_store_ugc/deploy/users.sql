-- Deploy book_store_ugc:users to pg

BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id           UUID PRIMARY KEY NOT NULL   DEFAULT uuid_generate_v4(),
    email        VARCHAR(128)     NOT NULL,
    password     VARCHAR(128)     NOT NULL,
    verified     BOOLEAN                     DEFAULT FALSE,
    first_name   VARCHAR(128)                DEFAULT '',
    last_name    VARCHAR(128)                DEFAULT '',
    phone_number VARCHAR(32)                 DEFAULT '',
    created_at   timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    last_login   timestamp without time zone DEFAULT timezone('UTC'::text, now())
);

CREATE UNIQUE INDEX users__email ON users (email);


COMMIT;
