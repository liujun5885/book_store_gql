-- Deploy book_store_assets:assets to pg

BEGIN;

CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- create books table
CREATE TABLE books
(
    id             uuid          DEFAULT uuid_generate_v4() NOT NULL primary key,
    safari_book_id VARCHAR(32)                              NOT NULL,
    status         smallint      DEFAULT 0                  NOT NULL,
    reviews        integer       DEFAULT 0                  NOT NULL,
    rating         integer       DEFAULT 0                  NOT NULL,
    popularity     integer       DEFAULT 0                  NOT NULL,
    report_score   integer       DEFAULT 0                  NOT NULL,
    pages          integer       DEFAULT 0                  NOT NULL,
    title          text                                     NOT NULL,
    description    text          DEFAULT ''                 NOT NULL,
    content        text          DEFAULT ''                 NOT NULL,
    source         VARCHAR(128)  DEFAULT ''                 NOT NULL,
    language       VARCHAR(32)   DEFAULT ''                 NOT NULL,
    tags           VARCHAR(32)[],
    url            VARCHAR(4096) DEFAULT ''                 NOT NULL,
    web_url        VARCHAR(4096) DEFAULT ''                 NOT NULL,
    created_at     timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    updated_at     timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    issued_at      timestamp without time zone
);

CREATE INDEX books__safari_book_id ON books USING btree (safari_book_id);
CREATE INDEX books__title ON books USING gin (title gin_trgm_ops);
CREATE INDEX books__reviews__rating ON books USING btree (reviews, rating);

-- create authors table
CREATE TABLE authors
(
    id          uuid        DEFAULT uuid_generate_v4() NOT NULL primary key,
    name        VARCHAR(256)                           NOT NULL,
    nationality VARCHAR(32) DEFAULT '',
    description text        DEFAULT '',
    created_at  timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    updated_at  timestamp without time zone DEFAULT timezone('UTC'::text, now())
);

CREATE INDEX authors__name ON authors USING btree (lower (name));

-- create publishers table
CREATE TABLE publishers
(
    id          uuid         DEFAULT uuid_generate_v4() NOT NULL primary key,
    name        VARCHAR(1024)                           NOT NULL,
    headquarter VARCHAR(128) DEFAULT '',
    description text         DEFAULT '',
    created_at  timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    updated_at  timestamp without time zone DEFAULT timezone('UTC'::text, now())
);

CREATE INDEX publisher__name ON publishers USING btree (lower (name));

-- create topics table
CREATE TABLE topics
(
    id         uuid DEFAULT uuid_generate_v4() NOT NULL primary key,
    name       VARCHAR(256)                    NOT NULL,
    slug       VARCHAR(256)                    NOT NULL,
    score      integer,
    created_at timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    updated_at timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    CONSTRAINT slug_unique UNIQUE (slug)
);
CREATE INDEX topic__name ON topics USING btree (lower (name));

-- create book author relation db
CREATE TABLE book_authors
(
    book_id    UUID NOT NULL REFERENCES books (id) ON DELETE CASCADE,
    author_id  UUID NOT NULL REFERENCES authors (id) ON DELETE CASCADE,
    created_at timestamp without time zone DEFAULT timezone('UTC'::text, now())
);

CREATE INDEX book_authors__book_id__author_id ON book_authors USING btree (book_id, author_id);

-- create book publisher relation db
CREATE TABLE book_publishers
(
    book_id      UUID NOT NULL REFERENCES books (id) ON DELETE CASCADE,
    publisher_id UUID NOT NULL REFERENCES publishers (id) ON DELETE CASCADE,
    created_at   timestamp without time zone DEFAULT timezone('UTC'::text, now())
);

CREATE INDEX book_publishers__book_id__publisher_id ON book_publishers USING btree (book_id, publisher_id);

-- create book topic relation db
CREATE TABLE book_topics
(
    book_id    UUID NOT NULL REFERENCES books (id) ON DELETE CASCADE,
    topic_id   UUID NOT NULL REFERENCES topics (id) ON DELETE CASCADE,
    created_at timestamp without time zone DEFAULT timezone('UTC'::text, now())
);

CREATE INDEX book_topics__book_id__topic_id ON book_topics USING btree (book_id, topic_id);

COMMIT;
