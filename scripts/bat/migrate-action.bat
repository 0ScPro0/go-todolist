@echo off
setlocal enabledelayedexpansion

if "%1"=="" (
    echo required parameter action is missing
    echo Usage: migrate-action.bat [up^|down]
    exit /b 1
)

if not exist ".env" (
    echo .env file not found
    exit /b 1
)

for /f "usebackq delims=" %%i in (".env") do set "%%i"

if "%DATABASE_POSTGRES_URL%"=="" (
    echo DATABASE_POSTGRES_URL is not set in .env
    exit /b 1
)

docker compose run --rm todolist-postgres-migrate -path /migrations -database "%DATABASE_POSTGRES_URL%" "%1"