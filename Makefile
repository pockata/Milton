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
API_RUN ?= $(DOCKER_COMPOSE) -f docker-compose.base.yml -f docker-compose.dev.yml run --rm api
UI_RUN ?= $(DOCKER_COMPOSE) -f docker-compose.base.yml -f docker-compose.dev.yml run --rm ui
GO_RUN ?= $(API_RUN) go
PRISMA:=$(API_RUN) ./scripts/prisma.sh
SQLBOILER:=$(API_RUN) ./scripts/sqlboiler.sh

---------------------------
---       General       ---
---------------------------

.DEFAULT_GOAL=help

.PHONY: help
help:
	@echo ""
	@echo "Run make dev"
	@echo ""

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

---------------------------
---       backend       ---
---------------------------

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

.PHONY: orm-generate
orm-generate:
	@echo -e "Generating ORM...\n"
	@$(SQLBOILER) sqlite3

.PHONY: mod-tidy
mod-tidy:
	$(GO_RUN) mod tidy
