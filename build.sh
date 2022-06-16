#!/bin/sh
#

export CGO_ENABLED=0

build() {
    GOOS=$1
    GOARCH=$2
    echo building for $1/$2
    go build -o build/$1-$2$3 main.go
}

####################################################################

build android arm64

build darwin amd64
build darwin arm64

build linux 386
build linux amd64
build linux arm64

build windows 386 .exe
build windows amd64 .exe
build windows arm64 .exe
