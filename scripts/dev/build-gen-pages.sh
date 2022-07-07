#!/bin/bash

go build -o gen-pages gen-pages.go
cp -f gen-pages dist/.