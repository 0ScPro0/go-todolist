SHELL := bash

# Define OS
ifeq ($(OS),Windows_NT)
    MIGRATE_ACTION = scripts\bat\migrate-action.bat
    MIGRATE_CREATE = scripts\bat\migrate-create.bat
    RM = del /Q
else
    MIGRATE_ACTION = ./scripts/sh/migrate-action.sh
    MIGRATE_CREATE = ./scripts/sh/migrate-create.sh
    RM = rm -f
endif

# Load .env
include .env
export

# ========= Commands =========

# Up postgres  container
env-up:
	docker compose up -d todolist-postgres

# Down postgres container
env-down:
	docker compose down -v todolist-postgres

# Create db migration
migrate-create:
	$(MIGRATE_CREATE) $(seq)

# Up db migration
migrate-up:
	$(MIGRATE_ACTION) up

# Down db migration
migrate-down:
	$(MIGRATE_ACTION) down

# Use db migration with custom action
migrate-action:
	$(MIGRATE_ACTION) $(action)

# Check migrate status
migrate-status:
	$(MIGRATE_ACTION) status

# Run backend application
run-server:
	@go mod tidy && go run cmd/server/main.go

.PHONY: env-up env-down migrate-create migrate-up migrate-down migrate-action run migrate-status migrate-force