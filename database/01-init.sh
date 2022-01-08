#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --PASSWORD "$POSTGRES_PASSWORD" <<-EOSQL
  CREATE DATABASE $DB_NAME;
EOSQL