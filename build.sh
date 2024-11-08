#!/usr/bin/env sh

set -e

GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/
GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o build/hw_arm
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/

[ "$1" = "c" ] || exit
cd build || exit

tar -czvf hw_Linux_aarch64.tar.gz hw_arm
tar -czvf hw_Linux_x86_64.tar.gz hw
zip -r hw_windows_x86_64.zip hw.exe


