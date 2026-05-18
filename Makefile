include .env
export

env-up:
	docker compose up todolist-postgres

env-down:
	docker compose down todolist-postgres