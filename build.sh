#!/usr/bin/env bash

PLATFORMS=("linux/amd64" "linux/arm64" "windows/amd64" "darwin/amd64")
APP_NAME="any_chat_id_bot"

mkdir -p build

for PLATFORM in "${PLATFORMS[@]}"; do
    GOOS=$(echo $PLATFORM | cut -d/ -f1)
    GOARCH=$(echo $PLATFORM | cut -d/ -f2)
    export GOOS
    export GOARCH
    echo "Building for ${GOOS}_${GOARCH}"
    if [ "$GOOS" == "windows" ]; then
        go build -o "build/${GOOS}_${GOARCH}/${APP_NAME}.exe"
    else
        go build -o "build/${GOOS}_${GOARCH}/${APP_NAME}"
    fi

    zip -r "build/${APP_NAME}_${GOOS}_${GOARCH}.zip" "build/${GOOS}_${GOARCH}/"
done
