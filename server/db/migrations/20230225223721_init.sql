-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users
(
    id         UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    email      TEXT        NOT NULL,
    password   TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

COMMENT ON TABLE users IS 'Users';
COMMENT ON COLUMN users.id IS 'ID';
COMMENT ON COLUMN users.email IS 'Email';
COMMENT ON COLUMN users.password IS 'Password';
COMMENT ON COLUMN users.created_at IS 'Create date';
COMMENT ON COLUMN users.updated_at IS 'Update date';

-- Создаем уникальный индекс для users, чтобы email был уникальным и чтобы быстрее делать поиск по email
CREATE UNIQUE INDEX users_email_uidx ON users USING btree (email);

CREATE TABLE tasks
(
    id         UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    user_id    UUID        NOT NULL,
    title      TEXT        NOT NULL DEFAULT '',
    status     INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

COMMENT ON TABLE tasks IS 'Tasks';
COMMENT ON COLUMN tasks.id IS 'ID';
COMMENT ON COLUMN tasks.user_id IS 'User ID';
COMMENT ON COLUMN tasks.title IS 'Title';
COMMENT ON COLUMN tasks.status IS 'Status';
COMMENT ON COLUMN tasks.created_at IS 'Create date';
COMMENT ON COLUMN tasks.updated_at IS 'Update date';

-- Создаем индекс для tasks, чтобы быстрее делать поиск по user_id
CREATE INDEX tasks_user_id_idx ON tasks USING btree (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
