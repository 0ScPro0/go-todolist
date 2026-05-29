CREATE SCHEMA IF NOT EXISTS gotodolist;

CREATE TABLE IF NOT EXISTS gotodolist.tasks (
    id           SERIAL                   PRIMARY KEY,
    version      BIGINT         NOT NULL  DEFAULT 1,
    name         VARCHAR(100)   NOT NULL  CHECK (char_length(name) BETWEEN 1 AND 100),
    description  VARCHAR(1000)            CHECK (char_length(description) BETWEEN 1 AND 1000),
    status       SMALLINT       NOT NULL  DEFAULT 0 CHECK (status IN (0, 1, 2)),
    created_at   TIMESTAMPTZ    NOT NULL  DEFAULT CURRENT_TIMESTAMP
)