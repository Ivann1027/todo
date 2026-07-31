#!/bin/sh
set -e

echo "Waiting for database..."
until pg_isready -h db -p 5432 -U ivan -d todo-db; do
  sleep 2
done

echo "Running migrations..."
export PGPASSWORD=1234
goose -dir ./migrations postgres "host=db user=ivan dbname=todo-db sslmode=disable" up

echo "Starting app..."
exec ./todo-api