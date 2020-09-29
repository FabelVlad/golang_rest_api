.PHONY: build
build:
	go build -v ./cmd/server
.PHONY: test
test:
	go test -v -race -timout 30s ./...
.DEFAULT_GOAL := build