#!/bin/sh

until psql -v ON_ERROR_STOP=1 "$POSTGRES_URI"; do
  >&2 echo "Waiting postgres container to start"
  sleep 1
done

printf "Creating database \n"
psql -v ON_ERROR_STOP=1 "$POSTGRES_URI" -f ./create-db.sql

migrate -source file://./migrations -database "$POSTGRES_URI_WITH_DB" up

go run cmd/api/main.go
