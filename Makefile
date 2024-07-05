.DEFAULT_GOAL := run

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

build: fmt
	go build -mod=vendor -o htmx_blog
.PHONY: build

run: fmt templ_gen
	air
.PHONY:run
