SHELL := bash

include .env
export

env-up:
	docker compose up -d todolist-postgres

env-down:
	docker compose down -v todolist-postgres

migrate-create:
	scripts\migrate-create.bat $(seq)

run:
	@go run cmd/server/main.go