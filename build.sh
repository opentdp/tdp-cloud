#!/bin/sh
#

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    echo building for $1/$2
    GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o build/tdp-cloud-$1-$2$3 main.go
}

####################################################################

build android arm64

build darwin amd64
build darwin arm64

build linux amd64
build linux arm64

build windows amd64 .exe
build windows arm64 .exe
