//go:build prod

package env

var folder = "/Users/Shared/.bbts"

var EtcHostsPath = "/etc/hosts"
var DBPath = SafeFile(folder, "db", "db")
var DaemonName = "com.betterblockedthansorry.bbts"
var FirstProgramPath = "/Applications/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon"
var ProgramPath = SafeFile(folder, "bbts_daemon")
var DownloadPath = SafePath(folder, "download")
