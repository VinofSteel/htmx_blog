.DEFAULT_GOAL := run

fmt:
	gofmt -w .
.PHONY:fmt

vendor:
	go mod vendor
.PHONY: vendor

build: fmt
	go build -mod=vendor -o htmx_blog
.PHONY: build

run: fmt
	air
.PHONY:run
