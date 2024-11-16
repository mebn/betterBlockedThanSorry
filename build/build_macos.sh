#!/bin/bash

# Determine the directory where the script resides and navigate to it
BUILD_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$BUILD_DIR"
cd .. # Navigate to the root folder

wails build -tags "prod"

# needs to be installed somewhere else to prevent user from uninstalling
go build -tags "prod" -o build/bin/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon cmd/daemon/bbts_daemon.go

cp build/darwin/launch_with_sudo.sh build/bin/BetterBlockedThanSorry.app/Contents/MacOS/launch_with_sudo.sh
