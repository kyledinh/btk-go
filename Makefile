ROOT := github.com/kyledinh/btk-go

export SHELL := /bin/bash

BUILD_DIR = ./build
OUTPUT_DIR = ./dist

# Current version of the project.
GIT_TAG ?= $(shell git describe --tags --always --dirty)

# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
GOROOT ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

# ACTIONS
build:
	@echo "## Building the binaries"
	GOOS=darwin GOARCH=amd64 go build -o dist/btk-cli-macos cmd/cli/main.go
	GOOS=linux GOARCH=386 go build -ldflags="-X 'btk-go/pkg/version.Version=$(GIT_TAG)'" -o dist/btk-cli-linux cmd/cli/main.go
	@echo "dist/"
	@ls dist

deploy:
    DEPLOY_CMD := $(shell cp ./dist/btk-cli-macos /Users/kyle/bin/)
	
check:
	@./scripts/dev/check.sh

lint: 
	@./scripts/dev/lint.sh

analyze:
	@./scripts/dev/lint.sh
	go vet -v cmd/...
	staticcheck github.com/kyledinh/btk-go/cmd/...

setup:
	@./scripts/dev/setup.sh

test:
	go test ./...
