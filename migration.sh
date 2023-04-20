#!/bin/sh

set -e

# Define your database configuration
export POSTGRES_HOST=127.0.0.1
export POSTGRES_PORT=5432
export POSTGRES_USER=demyst
export POSTGRES_PASSWORD=admin
export POSTGRES_DB=loan-db


docker run -v $PWD/migration:/migrations --network host migrate/migrate \
    -path=/migrations/ \
    -database postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable up