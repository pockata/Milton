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

DOCKER_COMPOSE = docker compose
CONTAINER_RUN ?= $(DOCKER_COMPOSE) -f docker-compose.base.yml -f docker-compose.dev.yml run --rm api
GO_RUN ?= $(CONTAINER_RUN) go
BOILER:=sqlboiler
SRV:=$(shell realpath "packages/server")
DB:=$(shell realpath "${SRV}/adapters/db")
PRISMA:=$(CONTAINER_RUN) ./scripts/prisma.sh

.DEFAULT_GOAL=help

.PHONY: help
help:
	@echo ""
	@echo "Run make dev"
	@echo ""

.PHONY: schema-sync
schema-sync: schema-push orm-generate
	@echo "Schema synced"

.PHONY: schema-push
schema-push: 
	@$(PRISMA) db push

.PHONY: schema-status
schema-status:
	@$(PRISMA) migrate status

.PHONY: schema-migrate
schema-migrate:
	@$(PRISMA) migrate dev

.PHONY: schema-format
schema-format:
	@$(PRISMA) format

.PHONY: boiler-check
boiler-check:
	@which sqlboiler sqlboiler-sqlite3 > /dev/null 2>&1 || (\
		echo "SQLBoiler not found... installing" && \
		go install github.com/volatiletech/sqlboiler/v4@latest && \
		go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest \
	)

.PHONY: orm-generate
orm-generate: boiler-check
	@echo -e "Generating ORM...\n"
	@cd "${DB}" && SQLITE3_DBNAME="${SRV}/${DB_FILE}" $(BOILER) sqlite3

.PHONY: mod-tidy
mod-tidy:
	$(GO_RUN) mod tidy

.PHONY: dev
dev:
	@echo -e "Starting dev env\n"
	@$(DOCKER_COMPOSE) \
		-f docker-compose.base.yml \
		-f docker-compose.dev.yml \
		up

.PHONY: prod
prod:
	@echo -e "Starting prod env\n"
	@$(DOCKER_COMPOSE) \
		-f docker-compose.base.yml \
		-f docker-compose.prod.yml \
		up --build

