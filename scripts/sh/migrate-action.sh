#!/bin/bash
set -a
source .env
set +a

if [ -z "$1" ]; then
    echo "required parameter action is missing"
    echo "Usage: $0 [up|down]"
    exit 1
fi

if [ -z "$DATABASE_POSTGRES_URL" ]; then
    echo "DATABASE_POSTGRES_URL is not set in .env"
    exit 1
fi

docker compose run --rm todolist-postgres-migrate -path /migrations -database "$DATABASE_POSTGRES_URL" "$1"