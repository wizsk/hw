#!/usr/bin/env sh

set -e

GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/
GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o build/hw_arm
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/
