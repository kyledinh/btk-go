ROOT := github.com/kyledinh/btk-go

export SHELL := /bin/bash

BUILD_DIR = ./build
OUTPUT_DIR = ./dist

# Current version of the project.
GITTAG ?= $(shell git describe --tags --always --dirty)
SEMVER ?= $(shell head -n 1 semver)

# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
GOROOT ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

# ACTIONS
.PHONY: build test webserver-up

analyze:
	@./scripts/dev/lint.sh
	go vet -v cmd/...
	staticcheck github.com/kyledinh/btk-go/cmd/...

bench: 
	@go test -bench=. github.com/kyledinh/btk-go/pkg/moxutil -benchmem -run=10000

build:
	@echo "## Building the binaries : $(SEMVER)-$(GITTAG)"
	GOOS=linux GOARCH=386 go build -ldflags "-X 'config/config.Version=$(SEMVER)-$(GITTAG)'" -o dist/btk-cli-linux cmd/cli/main.go
	go build -ldflags "-X 'github.com/kyledinh/btk-go/config.Version=$(SEMVER)-$(GITTAG)'" -o dist/btk-cli-mac cmd/cli/main.go
	@echo "dist/"
	@ls dist
	go build -ldflags "-X 'github.com/kyledinh/btk-go/config.Version=$(SEMVER)-$(GITTAG)'" -o dist/btk-http cmd/http-server/main.go

check:
	@./scripts/dev/check.sh

deploy:
	cp ./dist/btk-cli-mac /Users/kyle/bin/btk
	btk -v

gen-petstore:
	btk-cli -i=specs/petstore.1.0.0.yaml -gen=model
	@cp dist/gen.model.* internal/model/.

generate:
	go generate ./pkg/petstore/...

lint: 
	@./scripts/dev/lint.sh

setup:
	@./scripts/dev/setup.sh

test:
	go test ./...

webserver-up:
	@echo "Running webserver on http://localhost:8001"
	./dist/btk-http