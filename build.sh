#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    echo building for $1/$2
    target=build/tdp-cloud-$1-$2
    if [ x"$1" = x"windows" ]; then
        target="${target}.exe"
    fi
    GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o $target main.go
}

####################################################################

build android arm64

build darwin amd64
build darwin arm64

build freebsd 386
build freebsd amd64
build freebsd arm64

build linux 386
build linux amd64
build linux arm64
build linux mips
build linux mipsle
build linux mips64
build linux mips64le
build linux ppc64
build linux ppc64le
build linux s390x

build netbsd 386
build netbsd amd64
build netbsd arm64

build openbsd 386
build openbsd amd64
build openbsd arm64

build windows 386
build windows amd64
build windows arm64

####################################################################

for app in `ls build`; do
    gzip build/$app
done
