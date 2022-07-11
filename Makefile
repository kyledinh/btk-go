ROOT := github.com/kyledinh/btk-go

export SHELL := /bin/bash

BUILD_DIR = ./build
OUTPUT_DIR = ./dist

# Current version of the project.
GITTAG ?= $(shell git describe --tags --always --dirty)
SEMVAR ?= $(shell head -n 1 semvar)

# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
GOROOT ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

# ACTIONS
.PHONY: build test

analyze:
	@./scripts/dev/lint.sh
	go vet -v cmd/...
	staticcheck github.com/kyledinh/btk-go/cmd/...

build:
	@echo "## Building the binaries"
	GOOS=linux GOARCH=386 go build -ldflags "-X 'config/config.Version=$(SEMVAR)-$(GITTAG)'" -o dist/btk-cli-linux cmd/cli/main.go
	go build -ldflags "-X 'github.com/kyledinh/btk-go/config.Version=$(SEMVAR)-$(GITTAG)'" -o dist/btk-cli-mac cmd/cli/main.go
	@echo "dist/"
	@ls dist

check:
	@./scripts/dev/check.sh

deploy:
    DEPLOY_CMD := $(shell cp ./dist/btk-cli-macos /Users/kyle/bin/)
	
lint: 
	@./scripts/dev/lint.sh

setup:
	@./scripts/dev/setup.sh

test:
	go test ./...
