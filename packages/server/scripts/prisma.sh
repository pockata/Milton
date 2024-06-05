#!/usr/bin/env sh

cd "adapters/db" && PRISMA_DB_FILE="file:../${DB_FILE}" prisma-client-go $@
