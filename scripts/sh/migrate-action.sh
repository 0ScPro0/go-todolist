#!/bin/bash
set -a
source .env
set +a

if [ -z "$1" ]; then
    echo "required parameter action is missing"
    echo "Usage: $0 [up|down|status|version|force|drop]"
    exit 1
fi


if [ "$DATABASE_POSTGRES_HOST" = "localhost" ]; then
    echo "Warning: DATABASE_POSTGRES_HOST is localhost, overriding to todolist-postgres for Docker"
    DB_HOST="todolist-postgres"
else
    DB_HOST="$DATABASE_POSTGRES_HOST"
fi

MIGRATE_URL="postgres://${DATABASE_POSTGRES_USER}:${DATABASE_POSTGRES_PASSWORD}@${DB_HOST}:${DATABASE_POSTGRES_PORT}/${DATABASE_POSTGRES_DB}?sslmode=disable"

echo "Running migration action: $1"
echo "Database: ${DATABASE_POSTGRES_DB}"
echo "Host: ${DB_HOST}"
echo

docker compose run --rm todolist-postgres-migrate -path /migrations -database "$MIGRATE_URL" "$1"

if [ $? -ne 0 ]; then
    echo "Migration failed!"
    exit 1
fi

echo "Migration completed successfully!"