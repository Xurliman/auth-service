-- +goose Up
-- +goose StatementBegin
    CREATE TABLE user_sessions
    (
        id            uuid primary key unique not null default uuid_generate_v4(),
        user_id       uuid                    not null,
        session_token text                    not null,
        expires_at    timestamp,
        is_active     bool,
        created_at    timestamp                        default now(),

        CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE IF EXISTS user_sessions;
-- +goose StatementEnd
