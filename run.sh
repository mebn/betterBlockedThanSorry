#!/bin/bash

mkdir -p build/bin/
go build -o build/bin/bbts_daemon cmd/daemon/bbts_daemon.go
wails dev
