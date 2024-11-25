//go:build prod

package env

var folder = "/Users/Shared/.bbts"

var EtcHostsPath = "/etc/hosts"
var DBPath = safePath(folder, "db", "db")
var DaemonName = "com.betterblockedthansorry.bbts"
var FirstProgramPath = "/Applications/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon"
var ProgramPath = safePath(folder, "bbts_daemon")
