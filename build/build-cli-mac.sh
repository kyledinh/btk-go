#!/bin/bash

SEMVAR=$(cat ../semvar)
GITTAG=$(git describe --tags --always --dirty)
VERSION="$SEMVAR-$GITTAG"

echo "Building version: $VERSION"

go build -ldflags="-X 'github.com/kyledinh/btk-go/config.Version=$VERSION'" -o ../dist/btk-cli-macos ../cmd/cli/main.go
cp ../dist/btk-cli-macos ~/bin/btk-cli

# go tool nm <your binary> | grep <your variable>
# *[main][~/src/github.com/kyledinh/btk-go]$ go tool nm ./dist/btk-cli-mac | grep ersion
#  12c8d40 D github.com/kyledinh/btk-go/config.Version