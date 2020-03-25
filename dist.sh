#!/bin/bash
set -xeo pipefail

NAME=nifi_exporter
VERSION=$(git describe | sed 's/^v//')
BUILD_DIR=./dist
GOOS_LIST=(linux darwin windows)
GOARCH_LIST=(386 amd64)

mkdir -p $BUILD_DIR
cd $BUILD_DIR

for os in "${GOOS_LIST[@]}"; do
    for arch in "${GOARCH_LIST[@]}"; do
        out="${NAME}-${VERSION}.${os}-${arch}"
        out_bin="${out}/${NAME}"
        if [ "$os" = "windows" ]; then
            out_bin="${out_bin}.exe"
        fi

        CGO_ENABLED=0 GOOS="$os" GOARCH="$arch" go build -ldflags="-w -s" -o "$out_bin" ..
        cp ../LICENSE "${out}/LICENSE"

        if [ "$os" = "windows" ]; then
            zip -r "${out}.zip" "$out"
        else
            tar -czf "${out}.tar.gz" "$out"
        fi
    done
done
