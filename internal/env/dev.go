//go:build !prod

package env

import (
	"fmt"
)

var EtcHostsPath = "/etc/hosts"
var LogPath = "/tmp/bbts.log"
var DaemonName = "com.betterblockedthansorry.bbts"
var ProgramPath = fmt.Sprintf("%s/../../build/bin/bbts_daemon", getCurrentFilePath())
var DBPath = "/tmp/badger"
