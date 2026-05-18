@echo off
if "%1"=="" (
    echo required seq parameter is missing
    exit /b 1
)

docker compose run --rm todolist-postgres-migrate create -ext sql -dir /migrations -seq "%1"