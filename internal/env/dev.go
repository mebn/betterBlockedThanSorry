//go:build !prod

package env

import (
	"fmt"
)

var EtcHostsPath = "/etc/hosts"
var DBPath = "/tmp/badger"
var DaemonName = "com.betterblockedthansorry.bbts"
var ProgramPath = fmt.Sprintf("%s/../../build/bin/bbts_daemon", getCurrentFilePath())
