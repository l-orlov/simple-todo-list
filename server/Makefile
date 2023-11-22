LOCAL_DB_HOST:=127.0.0.1
LOCAL_DB_NAME:=todo_list_local
LOCAL_DB_PORT:=5432
LOCAL_DB_DSN:=host=$(LOCAL_DB_HOST) port=$(LOCAL_DB_PORT) dbname=$(LOCAL_DB_NAME) sslmode=disable

LOCAL_DB_HOST:=127.0.0.1
LOCAL_DB_NAME:=todo_list_local
LOCAL_DB_PORT:=5432
LOCAL_DB_DSN:=host=127.0.0.1 port=5432 dbname=todo_list_local sslmode=disable


db-create-migration:
	goose -dir db/migrations create $(name) sql

db-migrate:
	goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" up
	make db-gen-structure

db-migrate-down:
	goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" down
	make db-gen-structure

db-reset:
	psql -c "drop database if exists $(LOCAL_DB_NAME)"
	psql -c "create database $(LOCAL_DB_NAME)"
	goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" up
	make db-gen-structure

db-gen-structure:
	pg_dump "$(LOCAL_DB_DSN)" --schema-only --no-owner --no-privileges --no-tablespaces --no-security-labels --no-comments > db/structure.sql
