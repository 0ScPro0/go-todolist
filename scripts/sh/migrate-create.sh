#!/bin/bash
if [ -z "$1" ]; then
    echo "required parameter seq is missing"
    echo "Usage: $0 <migration_name>"
    exit 1
fi

docker compose run --rm todolist-postgres-migrate create -ext sql -dir /migrations -seq "$1"