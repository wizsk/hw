#!/usr/bin/env sh

set -ex

# build directory
bd="build"
rm -rf "$bd/"*

GGOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/
tar -czvf "$bd/hw_Linux_x86_64.tar.gz"  -C "$bd" "hw"
rm "$bd/hw"

GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/
tar -czvf "$bd/hw_Linux_aarch64.tar.gz" -C "$bd" "hw"
rm "$bd/hw"

GOOS=windows GOARCH=amd64 CGO_ENABLED=0  go build -ldflags "-s -w" -o build/
zip -j "$bd/hw_windows_x86_64.zip" "$bd/hw.exe"
rm "$bd/hw.exe"
