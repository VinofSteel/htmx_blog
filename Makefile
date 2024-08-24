.DEFAULT_GOAL := run

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# This here is done to allow goose to run migrations, since it won't work using db as the value in a development environment
ifeq ($(ENV),production)
    PGHOST_OVERRIDE=$(PGHOST)
else
    PGHOST_OVERRIDE=localhost
endif

PG_CONN_STRING=postgresql://$(PGUSER):$(PGPASSWORD)@$(PGHOST_OVERRIDE):$(PGPORT)/$(PGDATABASE)?sslmode=disable

# Templ
templ_gen:
	templ generate
.PHONY:templ_gen

# API
fmt:
	gofmt -w .
.PHONY:fmt

vendor:
	go mod vendor
.PHONY: vendor

build: vendor fmt templ_gen
	go build -mod=vendor -o templ_blog
.PHONY: build

run: fmt templ_gen
	air
.PHONY:run

test: fmt
	go test ./... -count=1
.PHONY: test

sqlc-gen: fmt 
	sqlc generate
.PHONY:sqlc

m-up:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" up
.PHONY: m-up

m-down:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" down
.PHONY: m-down

m-status:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" status
.PHONY: m-status
