include .env
export

env-up:
	docker compose up -d todolist-postgres

env-down:
	docker compose down -v todolist-postgres