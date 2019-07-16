#!/bin/bash

export RELEASEDIR=./release
export VERSION=0.2.9

CGO_ENABLED=1 GOOS=linux  go build  -o $RELEASEDIR/build/usr/sbin/ovpmd ./cmd/ovpmd
CGO_ENABLED=1 GOOS=linux  go build  -o $RELEASEDIR/build/usr/bin/ovpm   ./cmd/ovpm

tar cvzf $RELEASEDIR/ovpm-v$VERSION.tar.gz -C $RELEASEDIR/build .

echo "Build done."
