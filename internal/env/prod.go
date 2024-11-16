//go:build prod

package env

var EtcHostsPath = "/etc/hosts"
var LogPath = "/tmp/bbts.log"
var DaemonName = "com.betterblockedthansorry.bbts"

// needs to be installed somewhere else to prevent user from uninstalling
var ProgramPath = "/Applications/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon"
