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

releases_url=https://api.github.com/repos/opentdp/tdp-cloud-ui/releases/latest
download_url=`wget -qO- $releases_url | grep releases/download | cut -f4 -d "\""`

wget -O cloud-ui.tar.gz $download_url
tar xvf cloud-ui.tar.gz --strip-components 2 -C front

####################################################################

RUN_NUMBER=${GITHUB_RUN_NUMBER:-0}

last_tag=`git tag | sort -V | tail -n 1`
prev_tag=`git tag | sort -V | tail -n 2 | head -n 1`
git log $prev_tag..$last_tag --pretty=format:"%s" | grep -v "^release" | sed 's/^/- /' | sort > RELEASE.md

version=`echo $last_tag | sed 's/^v//'`
sed -i "s/^const Version = \".*\"/const Version = \"$version\"/" cmd/args/build.go

build_version=$((`grep -oP 'BuildVersion = "\K\d+' cmd/args/build.go` + $RUN_NUMBER))
sed -i "s/^const BuildVersion = \".*\"/const BuildVersion = \"$build_version\"/" cmd/args/build.go

echo "build info - tag: $last_tag, version: $version, build: $build_version"

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
build linux ppc64le
build linux s390x

build netbsd amd64

build openbsd amd64
build openbsd arm64

build windows amd64
build windows arm64

####################################################################

for app in `ls build`; do
    gzip build/$app
done
