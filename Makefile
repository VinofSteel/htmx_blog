.DEFAULT_GOAL := run

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

PG_CONN_STRING=postgresql://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(PGDATABASE)?sslmode=disable

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

build: fmt templ_gen
	go build -mod=vendor -o htmx_blog
.PHONY: build

run: fmt templ_gen
	air
.PHONY:run

m-up:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" up
.PHONY: m-up

m-down:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" down
.PHONY: m-down

m-status:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" status
.PHONY: m-status
