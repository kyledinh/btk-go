#!/bin/bash

SEMVAR=$(cat ../semvar)
GITTAG=$(git describe --tags --always --dirty)
VERSION="$SEMVAR-$GITTAG"

echo "Building version: $VERSION"

go build -ldflags="-X 'btk-go/pkg/version.Version=$VERSION'" -o ../dist/btk-cli-macos ../cmd/cli/main.go
cp ../dist/btk-cli-macos ~/bin/btk-cli