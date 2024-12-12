-- +goose Up
-- +goose StatementBegin
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE IF NOT EXISTS users
    (
        id                  uuid primary key unique not null default uuid_generate_v4(),
        username            varchar(120)            not null,
        name                varchar(255)            not null,
        email               varchar(120)     unique not null,
        password            varchar(120)            not null,
        created_at          timestamp                         default now(),
        updated_at          timestamp,
        deleted_at          timestamp
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE IF EXISTS users;
-- +goose StatementEnd
