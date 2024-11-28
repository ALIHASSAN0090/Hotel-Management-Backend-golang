#!/bin/bash
set -e

until pg_isready -U "$POSTGRES_USER" -d "$POSTGRES_DB"; do
    echo "Waiting for PostgreSQL to be ready..."
    sleep 2
done

echo "Running migrations..."

if ! go mod tidy; then
    echo "Error running go mod tidy"
    exit 1
fi


if ! go run db_migrator/main.go db_migrator/migrations.go; then
    echo "Error running migrations"
    exit 1
fi

echo "Migrations completed!"
