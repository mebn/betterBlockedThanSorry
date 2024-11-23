//go:build prod

package env

var EtcHostsPath = "/etc/hosts"
var DBPath = "/tmp/badger"
var DaemonName = "com.betterblockedthansorry.bbts"

// TODO: needs to be installed somewhere else to prevent user from uninstalling
var ProgramPath = "/Applications/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon"
