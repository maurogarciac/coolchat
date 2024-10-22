PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})
GO_LINT=golangci/golangci-lint

.PHONY: lint
## lint: run linters and static analysis tools to detect potential issues, bugs, and code style violations in Go codebases
lint:
	go mod vendor
	docker run -t --rm -v $(shell pwd):/app -w /app ${GO_LINT} golangci-lint run --fix -v

## install: install dependencies for project
.PHONY: install
install:
	go install cmd/frontend/

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## debug: run the server in debug mode
.PHONY: debug
debug:
	go run cmd/main.go

## start: run the server locally
.PHONY: start
start:
	go run cmd/main.go

# build: build a binary
.PHONY: build
build:
	go build -o bin/frontend ./cmd

## startd: start docker container for server
.PHONY: startd
startd:
	docker run cc-frontend:latest

## binux: build docker container with app binary
.PHONY: binux
binux:
	@docker build \
	-t cc-frontend \
	--platform local .

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
