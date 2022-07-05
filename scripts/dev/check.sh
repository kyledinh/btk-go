#!/bin/bash

REPO="btk-go"
THISDIR=$(pwd)
REPO_DIR=${THISDIR/$REPO*/$REPO}
VERSION=$(cat $REPO_DIR/semvar) 

echo "Checking development environment"
echo "REPO_DIR: $REPO_DIR"
echo "SEMVAR: $VERSION"
