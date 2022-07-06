#!/bin/bash

REPO="btk-go"
THISDIR=$(pwd)
REPO_DIR=${THISDIR/$REPO*/$REPO}

docker run --rm \
    -v "$REPO_DIR:$REPO_DIR" \
    -w "$REPO_DIR" \
    --entrypoint golangci-lint \
    golangci/golangci-lint:v1.41.1 \
    run -v --new --timeout 5m