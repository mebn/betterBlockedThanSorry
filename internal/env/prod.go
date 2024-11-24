//go:build prod

package env

var EtcHostsPath = "/etc/hosts"
var DBPath = "/tmp/.bbtsdb/bbtsdb.db"
var DaemonName = "com.betterblockedthansorry.bbts"

// TODO: needs to be installed somewhere else to prevent user from uninstalling
var ProgramPath = "/Applications/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon"
