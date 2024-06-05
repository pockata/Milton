#!/usr/bin/env sh

cd "adapters/db" && SQLITE3_DBNAME="../../${DB_FILE}" sqlboiler $@

