#!/bin/bash

mkdir -p build/bin/
go build -o build/bin/bbts_daemon cmd/daemon/main.go
cd cmd/main
wails dev
