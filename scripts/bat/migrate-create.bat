@echo off
if "%1"=="" (
    echo required parameter seq is missing
    echo Usage: migrate-create.bat ^<migration_name^>
    exit /b 1
)

docker compose run --rm todolist-postgres-migrate create -ext sql -dir /migrations -seq "%1"