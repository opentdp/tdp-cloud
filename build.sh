#!/bin/sh
#

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    echo building for $1/$2
    out=build/tdp-cloud-$1-$2$3
    GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o $out main.go
}

####################################################################

build android arm64

build darwin amd64
build darwin arm64

build linux amd64
build linux arm64

build windows amd64 .exe
build windows arm64 .exe

####################################################################

if type apt >/dev/null 2>&1; then
    sudo apt install -y upx-ucl
fi

if type apk >/dev/null 2>&1; then
    apk add upx
fi

if type upx >/dev/null 2>&1; then
    cd build
    upx --best `ls .`
    cd ..
fi
