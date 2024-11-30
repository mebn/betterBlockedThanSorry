#!/bin/bash

mkdir -p build/bin/

# blocker daemon
go build -o build/bin/bbts_daemonDEV cmd/daemon/main.go

# updater agent
go build -o build/bin/bbts_updaterDEV cmd/updaterAgent/main.go

# dist folder can't be empty
mkdir -p frontend/dist
touch frontend/dist/.tempfile

cd cmd/main
wails dev
