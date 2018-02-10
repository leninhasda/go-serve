#!/bin/bash

echo "Building for mac..."
GOARCH="amd64"	GOOS="darwin" CGO_ENABLED=0 go build -o go-serve-mac

echo "Building for linux..."
GOARCH="amd64" GOOS="linux" CGO_ENABLED=0 go build -o go-serve-linux

echo "Building for windows..."
GOARCH="amd64" GOOS="windows" CGO_ENABLED=0 go build -o go-serve.exe
