//go:build !prod

package env

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var EtcHostsPath = "/etc/hosts"
var LogPath = "/tmp/bbts.log"
var DaemonName = "com.betterblockedthansorry.bbts"

// var ProgramPath = "/Users/mebn/go/bin/bbts_daemon"
// var ProgramPath = "/Users/mebn/marcus/code/betterBlockedThanSorry/build/cmd"
var ProgramPath = fmt.Sprintf("%s/../../build/bin/bbts_daemon", getCurrentFilePath())

func getCurrentFilePath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Unable to get current file info")
	}
	return filepath.Dir(file)
}
