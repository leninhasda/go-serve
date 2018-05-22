#!/bin/bash

# echo "Building for mac..."
# GOARCH="amd64"	GOOS="darwin" CGO_ENABLED=0 go build -o go-serve-mac

# echo "Building for linux..."
# GOARCH="amd64" GOOS="linux" CGO_ENABLED=0 go build -o go-serve-linux

# echo "Building for windows..."
# GOARCH="amd64" GOOS="windows" CGO_ENABLED=0 go build -o go-serve.exe


# Build 386 amd64 binaries
OS_PLATFORM_ARG=(linux darwin windows freebsd openbsd)
OS_ARCH_ARG=(386 amd64)
for OS in ${OS_PLATFORM_ARG[@]}; do
  for ARCH in ${OS_ARCH_ARG[@]}; do
    echo "Building binary for $OS/$ARCH..."
    GOARCH=$ARCH GOOS=$OS CGO_ENABLED=0 go build -o "dist/go-serve_$OS-$ARCH" .
  done
done
