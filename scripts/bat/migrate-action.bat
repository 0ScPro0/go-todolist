@echo off
setlocal enabledelayedexpansion

if "%1"=="" (
    echo required parameter action is missing
    echo Usage: migrate-action.bat [up^|down^|status^|version^|force^|drop]
    exit /b 1
)

if not exist ".env" (
    echo .env file not found
    exit /b 1
)

for /f "usebackq delims=" %%i in (".env") do set "%%i"

set "MIGRATE_URL=postgres://%DATABASE_POSTGRES_USER%:%DATABASE_POSTGRES_PASSWORD%@todolist-postgres:%DATABASE_POSTGRES_PORT%/%DATABASE_POSTGRES_DB%?sslmode=disable"

echo Running migration action: %1
echo Database: %DATABASE_POSTGRES_DB%
echo URL: %MIGRATE_URL%
echo.

docker compose run --rm todolist-postgres-migrate -path /migrations -database "%MIGRATE_URL%" %1

if errorlevel 1 (
    echo Migration failed with error code: !errorlevel!
    exit /b !errorlevel!
)

echo Migration completed successfully!