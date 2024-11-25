# betterBlockedThanSorry

Block access to websites for a certain time period.

## Dev

Run:

```sh
# main (gui)
./run.sh

# daemon
go run cmd/daemon/bbts_daemon.go
```

### How to run tests

Run:

```sh
# run all tests
go test ./...
```

## Prod

### Build

```sh
# build for MacOS
./build/build_macos.sh
```
