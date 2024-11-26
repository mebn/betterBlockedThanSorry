//go:build !prod

package env

var folder = "/Users/Shared/.bbtsDEV"

var EtcHostsPath = "/etc/hosts"
var DBPath = SafeFile(folder, "db", "db")
var DaemonName = "com.betterblockedthansorry.bbtsDEV"
var FirstProgramPath = SafeFile(currentDir(), "../", "../", "build", "bin", "bbts_daemonDEV")
var ProgramPath = SafeFile(folder, "bbts_daemonDEV")
var DownloadPath = SafePath(folder, "download")
