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

## start: run the server
.PHONY: start
start:
	go run cmd/frontend/main.go

# build: build a binary
.PHONY: build
build:
	go build cmd/frontend/ -o ./bin

## startd: start docker container for server
.PHONY: startd
startd:
	docker run frontend:latest

## binux: build docker container with app binary
.PHONY: binux
binux:
	@docker build \
	-t frontend \
	--platform local .

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
