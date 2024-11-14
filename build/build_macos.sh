#!/bin/bash

BUILD_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$BUILD_DIR"
cd .. # back to root folder

wails build
go build -o build/bin/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon cmd/daemon/bbts_daemon.go
cp build/darwin/launch_with_sudo.sh build/bin/BetterBlockedThanSorry.app/Contents/MacOS/launch_with_sudo.sh
