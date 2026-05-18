include .env
export

export PROJECT_ROOT=$(pwd)

env-up:
	docker compose up todolist-postgres

env-down:
	docker compose down todolist-postgres