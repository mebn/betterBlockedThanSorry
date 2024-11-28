package main

import (
	"fmt"
	"os"

	"github.com/mebn/betterBlockedThanSorry/internal/daemon"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
	"github.com/mebn/betterBlockedThanSorry/internal/updater"
)

func main() {
	var err error

	file, _ := os.OpenFile(env.SafeFile(env.BaseFolder, "bbtsupdater.log"),
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	file.WriteString("[INFO] Agent started new.\n")

	appUpdater, err := updater.NewUpdater()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to create Updater: %s\n", err))
		stop(file)
	}

	binaryPath, err := appUpdater.DownloadLatestBinary()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to download latest binary: %s\n", err))
		stop(file)
	}

	err = appUpdater.ReplaceProgram(binaryPath, env.BinaryPath)
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to move the binary: %s\n", err))
		stop(file)
	}

	err = appUpdater.RelaunchProgram(env.BinaryPath)
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to reopen the binary: %s\n", err))
		stop(file)
	}

	stop(file)
}

func stop(file *os.File) {
	agent := daemon.NewAgent(env.UpdaterAgentName, env.UpdateProgramPath)
	err := agent.Stop()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Stopping agent failed. err: %s\n", err))
	}
}
