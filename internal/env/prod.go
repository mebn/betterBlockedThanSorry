//go:build prod

package env

var BaseFolder = "/Users/Shared/.bbts"

var EtcHostsPath = "/etc/hosts"
var DBPath = SafeFile(BaseFolder, "db", "db")
var BinaryPath = "/Applications/BetterBlockedThanSorry.app"

// blocker daemon
var DaemonName = "com.betterblockedthansorry.bbts"
var FirstProgramPath = "/Applications/BetterBlockedThanSorry.app/Contents/MacOS/bbts_daemon"
var ProgramPath = SafeFile(BaseFolder, "bbts_daemon")

// updater agent
var UpdaterAgentName = "com.betterblockedthansorry.bbtsUpdater"
var UpdateProgramPath = "/Applications/BetterBlockedThanSorry.app/Contents/MacOS/bbts_updater"
var DownloadPath = SafePath(BaseFolder, "download")
var SkipUpdate = false
