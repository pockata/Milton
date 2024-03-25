# Load .env file
# https://lithic.tech/blog/2020-05/makefile-dot-env
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

PROJECT_ROOT:=$(shell pwd)

# VERSION is already defined when running in CI
ifndef VERSION
	VERSION := $(shell git rev-parse --short HEAD)
endif

PRISMA:=PRISMA_DB_FILE="file:../${DB_FILE}" prisma-client-go
BOILER:=sqlboiler
SRV:=$(shell realpath "packages/server")

schema-sync: schema-push orm-generate
	@echo "Schema synced"

prisma-check:
	@which prisma-client-go > /dev/null 2>&1 || (\
		echo "Prisma not found... installing" && \
		go install github.com/prisma/prisma-client-go@latest \
	)

schema-push: prisma-check
	@cd "${SRV}" && $(PRISMA) db push

schema-status: prisma-check
	@cd "${SRV}" && $(PRISMA) migrate status

schema-migrate: prisma-check
	@cd "${SRV}" && $(PRISMA) migrate dev

schema-format: prisma-check
	@cd "${SRV}" && $(PRISMA) format

boiler-check:
	@which sqlboiler sqlboiler-sqlite3 > /dev/null 2>&1 || (\
		echo "SQLBoiler not found... installing" && \
		go install github.com/volatiletech/sqlboiler/v4@latest && \
		go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest \
	)

orm-generate: boiler-check
	@echo -e "Generating ORM...\n"
	@cd "${SRV}" && SQLITE3_DBNAME="${SRV}/${DB_FILE}" $(BOILER) sqlite3

build-server:
	@cd "${SRV}" && go build -ldflags "-X milton.Build=${VERSION}" -o ./bin/milton ./cmd/api/milton.go

run-server: build-server
	@cd "${SRV}" && ./bin/milton

go-tidy:
	@cd "${SRV}" && go mod tidy

server-watcher-check:
	@which gow > /dev/null 2>&1 || (\
		echo "GOW not found... installing" && \
		go install github.com/mitranim/gow@latest \
	)

dev: server-watcher-check
	@cd "${SRV}" && gow -c run ./cmd/api/milton.go

