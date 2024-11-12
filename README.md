# betterBlockedThanSorry

Block access to websites for a certain time period.

## How to run

Run:

```sh
# main (gui)
wails dev

# daemon
go run cmd/daemon/bbts_daemon.go
```

### How to run tests

Run:

```sh
# run all tests
go test ./...
```

## Binary with privileges (MacOS)

```sh
# File: x.app/Contents/MacOS/launch_with_sudo.sh
#!/bin/bash
osascript -e 'do shell script "/Applications/myproject.app/Contents/MacOS/myproject" with administrator privileges'

# File: x.app/Info.plist
<key>CFBundleExecutable</key>
<string>launch_with_sudo.sh</string>
```

Also remember to put `bbts_daemon` in `x.app/Contents/MacOS/`.
