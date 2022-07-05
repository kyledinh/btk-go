ROOT := github.com/kyledinh/btk-go

export SHELL := /bin/bash

BUILD_DIR = ./build
OUTPUT_DIR = ./dist

# Current version of the project.
VERSION ?= $(shell git describe --tags --always --dirty)

# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

# ACTIONS
build:
	@echo "## Building the binaries"
	go build -o dist/btk-cli cmd/cli/main.go
	@echo "dist/"
	@ls dist

check:
	@./scripts/dev/check.sh
	
# more info about `GOGC` env: https://github.com/golangci/golangci-lint#memory-usage-of-golangci-lint
lint: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) run

$(GOLANGCI_LINT):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.23.6
