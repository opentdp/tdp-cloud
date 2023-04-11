#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    echo building for $1/$2
    out=build/tdp-cloud-$1-$2$3
    GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o $out main.go
}

####################################################################

releases_url=https://api.github.com/repos/open-tdp/tdp-cloud-ui/releases/latest
download_url=`wget -qO- $releases_url | grep releases/download | cut -f4 -d "\""`

wget -O cloud-ui.tar.gz $download_url
tar xvf cloud-ui.tar.gz --strip-components 2 -C front

####################################################################

build android arm64

build darwin amd64
build darwin arm64

build linux amd64
build linux arm64

build windows amd64 .exe
build windows arm64 .exe

####################################################################

for app in `ls build`; do
    gzip build/$app
done
