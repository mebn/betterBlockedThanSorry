package main

import (
	"fmt"
	"os"

	"github.com/mebn/betterBlockedThanSorry/internal/env"
	"github.com/mebn/betterBlockedThanSorry/internal/service"
	"github.com/mebn/betterBlockedThanSorry/internal/updater"
)

func main() {
	var err error

	file, _ := os.OpenFile(env.SafeFile(env.BaseFolder, "bbtsupdater.log"),
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if len(os.Args) < 2 {
		stop(file)
	}

	currentVersion := os.Args[1]

	file.WriteString("[INFO] Agent started new.\n")

	appUpdater, err := updater.NewUpdater()
	handleError(file, err, "[ERR] Failed to create Updater")

	if appUpdater.UpToDate(currentVersion) {
		stop(file)
	}

	binaryPath, err := appUpdater.DownloadLatestBinary()
	handleError(file, err, "[ERR] Failed to download latest binary")

	err = appUpdater.ReplaceProgram(binaryPath, env.BinaryPath)
	handleError(file, err, "[ERR] Failed to move the binary")

	err = appUpdater.RelaunchProgram(env.BinaryPath)
	handleError(file, err, "[ERR] Failed to reopen the binary")

	stop(file)
}

func stop(file *os.File) {
	agent := service.NewBackgroundService(env.UpdaterAgentName, env.UpdateProgramPath, service.Agent)
	err := agent.Stop()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Stopping agent failed. err: %s\n", err))
	}
}

func handleError(file *os.File, err error, msg string) {
	if err != nil {
		file.WriteString(fmt.Sprintf("%s: %s\n", msg, err))
		stop(file)
	}
}
