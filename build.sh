#!/bin/bash

BUILD_DIR=bin
NAME=$(basename $PWD)

mkdir -p $BUILD_DIR

for GOOS in linux; do
  for GOARCH in 386 amd64; do
    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -v -o $BUILD_DIR/$NAME-$GOOS-$GOARCH
    mv ${BUILD_DIR}/current-${GOOS}-${GOARCH} ${BUILD_DIR}/aws-param-inject-${GOOS}-${GOARCH}
  done
done
