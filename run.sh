#!/bin/bash

mkdir -p build/bin/

# blocker daemon
go build -o build/bin/bbts_daemonDEV cmd/daemon/main.go

# updater agent
go build -o build/bin/bbts_updaterDEV cmd/updaterAgent/main.go

cd cmd/main
wails dev
