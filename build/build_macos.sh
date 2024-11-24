#!/bin/bash

# determine the directory where the script resides and navigate to it
BUILD_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$BUILD_DIR"

# navigate to where wailts.json is
cd ../cmd/main

wails build -tags "prod"

# go back to root
cd ../../

# TODO: needs to be installed somewhere else to prevent user from uninstalling
go build -tags "prod" -o build/bin/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon cmd/daemon/main.go

cp build/darwin/launch_with_sudo.sh build/bin/BetterBlockedThanSorry.app/Contents/MacOS/launch_with_sudo.sh
