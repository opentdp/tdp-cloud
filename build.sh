#!/bin/sh
#

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    echo building for $1/$2

    out=build/tdp-cloud-$1-$2$3
    GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o $out main.go

    if type upx >/dev/null 2>&1; then
        upx --best $out
    fi
}

####################################################################

build android arm64

build darwin amd64
build darwin arm64

build linux amd64
build linux arm64

build windows amd64 .exe
build windows arm64 .exe
