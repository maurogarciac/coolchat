PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})
GO_LINT=golangci/golangci-lint

.PHONY: lint
## lint: run linters and static analysis tools to detect potential issues, bugs, and code style violations in Go codebases
lint:
	go mod vendor
	docker run -t --rm -v $(shell pwd):/app -w /app ${GO_LINT} golangci-lint run --fix -v

## init: initialize project (make init module=github.com/user/project)
.PHONY: install
install:
	go install cmd/backend/

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## start: run the server
.PHONY: start
start:
	go run cmd/backend/main.go

all: help

# build: build a binary
.PHONY: build
build:
	go build cmd/backend/ -o ./bin


## binux: build docker container with app binary
.PHONY: binux
binux:
	@docker build \
	-t backend \
	--platform local .

all: help

## startd: start the docker container for server
.PHONY: startd
startd:
	docker run backend:latest

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
