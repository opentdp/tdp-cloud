#!/bin/sh
#

WKDIR=`dirname $GITHUB_WORKSPACE`

# Create workspace

mkdir -p $WKDIR
cd $WKDIR

# Download source code

if [ ! -d tdp-cloud ]; then
    git clone https://github.com/tdp-resource/tdp-cloud.git
fi

if [ ! -d tdp-cloud-ui ]; then
    git clone https://github.com/tdp-resource/tdp-cloud-ui.git
fi

chmod +x $WKDIR/*/build.sh

# Compile front-end components

cd $WKDIR/tdp-cloud-ui
npm i && ./build.sh

cp -av $WKDIR/tdp-cloud-ui/build/* $WKDIR/tdp-cloud/front/

# Compile backend components

cd $WKDIR/tdp-cloud
go mod tidy && ./build.sh

for app in `ls build`; do
    gzip build/$app
done
