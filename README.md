# betterBlockedThanSorry
Block access to websites for a certain time period.

## How to run
Run:
```sh
go run cmd/main/main.go
```

## My plan
My plan is to have `main/main.go` craete a systemd/launchd service. This service will start the daemon located in `daemon/main.go`. This daemon will make sure that the blocklist of websites is not tempered with etc.
