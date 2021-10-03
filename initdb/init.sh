#!/bin/bash
set -e
echo DB NAME IS $POSTGRES_DB

export PGDATABASE=$POSTGRES_DB
export PGUSER=$POSTGRES_USER
export PGPASSWORD=$POSTGRES_PASSWORD

psql -v ON_ERROR_STOP=1 <<-EOSQL
    CREATE TABLE IF NOT EXISTS logs (
        ts timestamptz PRIMARY KEY,
        msg text
    );
EOSQL
