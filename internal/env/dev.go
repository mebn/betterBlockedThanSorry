//go:build !prod

package env

var BaseFolder = "/Users/Shared/.bbtsDEV"

var EtcHostsPath = "/etc/hosts"
var DBPath = SafeFile(BaseFolder, "db", "db")
var DownloadPath = SafePath(BaseFolder, "download")
var BinaryPath = "/Applications/BetterBlockedThanSorry.app"

// blocker daemon
var DaemonName = "com.betterblockedthansorry.bbtsDEV"
var FirstProgramPath = SafeFile(currentDir(), "../", "../", "build", "bin", "bbts_daemonDEV")
var ProgramPath = SafeFile(BaseFolder, "bbts_daemonDEV")

// updater agent
var UpdaterAgentName = "com.betterblockedthansorry.bbtsUpdaterDEV"
var UpdateProgramPath = SafeFile(currentDir(), "../", "../", "build", "bin", "bbts_updaterDEV")
var SkipUpdate = true
