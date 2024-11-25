#!/bin/bash

# determine the directory where the script resides and navigate to it
BUILD_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$BUILD_DIR"

# navigate to where wailts.json is
cd ../cmd/main

wails build -tags "prod"

# go back to root
cd ../../

go build -tags "prod" -o build/bin/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon cmd/daemon/main.go
# go build -tags "prod" -o /Users/Shared/.bbtsDEV/bbts_daemon cmd/daemon/main.go
