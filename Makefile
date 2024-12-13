MIGRATION_PATH := database/migrations
CONFIG_PATH := ./config.yaml

#yq should be installed in your OS, see https://github.com/mikefarah/yq
DB_CONNECTION := $(shell yq '.database.connection' $(CONFIG_PATH))
DB_HOST := $(shell yq '.database.host' $(CONFIG_PATH))
DB_NAME := $(shell yq '.database.name' $(CONFIG_PATH))
DB_USER := $(shell yq '.database.user' $(CONFIG_PATH))
DB_PASSWORD := $(shell yq '.database.password' $(CONFIG_PATH))

# database connection string (DSN)
DSN := 'host=${DB_HOST} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD}'

.PHONY: check-yq
check-yq:
	@if ! command -v yq &>/dev/null; then \
		echo "Error: yq is not installed. See https://github.com/mikefarah/yq"; \
		exit 1; \
	fi

.PHONY: migrate-up
migrate-up: check-yq
	goose --dir ${MIGRATION_PATH} ${DB_CONNECTION} ${DSN} up

.PHONY: migrate-down
migrate-down: check-yq
	goose --dir ${MIGRATION_PATH} ${DB_CONNECTION} ${DSN} down-to 0

.PHONY: migration
migration: check-yq
	@if [ -z "$(name)" ]; then \
    		echo "Usage: make migration name=<migration_file_name>"; \
    		exit 1; \
	fi
	goose --dir ${MIGRATION_PATH} ${DB_CONNECTION} ${DSN} create $(name) sql
