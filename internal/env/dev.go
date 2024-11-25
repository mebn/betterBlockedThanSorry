//go:build !prod

package env

var folder = "/Users/Shared/.bbtsDEV"

var EtcHostsPath = "/etc/hosts"
var DBPath = safePath(folder, "db", "db")
var DaemonName = "com.betterblockedthansorry.bbtsDEV"
var FirstProgramPath = safePath(currentDir(), "../", "../", "build", "bin", "bbts_daemonDEV")
var ProgramPath = safePath(folder, "bbts_daemonDEV")
