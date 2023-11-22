-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE OR REPLACE FUNCTION set_updated_at_column() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now() AT TIME ZONE 'utc';
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE users
(
    id         UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    name       TEXT        NOT NULL DEFAULT '',
    email      TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

COMMENT ON TABLE users IS 'Users';
COMMENT ON COLUMN users.id IS 'ID';
COMMENT ON COLUMN users.name IS 'Name';
COMMENT ON COLUMN users.email IS 'Email';
COMMENT ON COLUMN users.created_at IS 'Create date';
COMMENT ON COLUMN users.updated_at IS 'Update date';

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_column();

CREATE TABLE tasks
(
    id         UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    title      TEXT        NOT NULL DEFAULT '',
    status     INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

COMMENT ON TABLE tasks IS 'Tasks';
COMMENT ON COLUMN tasks.id IS 'ID';
COMMENT ON COLUMN tasks.title IS 'Title';
COMMENT ON COLUMN tasks.status IS 'Status';
COMMENT ON COLUMN tasks.created_at IS 'Create date';
COMMENT ON COLUMN tasks.updated_at IS 'Update date';

CREATE TRIGGER update_tasks_updated_at
    BEFORE UPDATE
    ON tasks
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
