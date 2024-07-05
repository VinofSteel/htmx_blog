.DEFAULT_GOAL := run

fmt:
	gofmt -w .
.PHONY:fmt

build: fmt
	go build && ./htmx_blog
.PHONY:build

run: fmt
	air
.PHONY:run
