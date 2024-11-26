package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/daemon"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

func main() {
	file, _ := os.OpenFile(env.SafeFile(env.BaseFolder, "bbtsupdater.log"),
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	file.WriteString("[INFO] Agent started new.\n")

	// updater, err := updater.NewUpdater()
	// if err != nil {
	// 	file.WriteString(fmt.Sprintf("[ERR] Failed to create Updater: %s\n", err))
	// 	stop(file)
	// }

	// binaryPath, err := updater.DownloadLatestBinary()
	// if err != nil {
	// 	file.WriteString(fmt.Sprintf("[ERR] Failed to download latest binary: %s\n", err))
	// 	stop(file)
	// }

	// err = updater.ReplaceProgram(binaryPath, env.BinaryPath)
	// if err != nil {
	// 	file.WriteString(fmt.Sprintf("[ERR] Failed to move the binary: %s\n", err))
	// 	stop(file)
	// }

	time.Sleep(5 * time.Second)
	file.WriteString("[INFO] after sleep.\n")
	stop(file)
	file.WriteString("[INFO] after stop.\n")
}

func stop(file *os.File) {
	agent := daemon.NewAgent(env.UpdaterAgentName, env.UpdateProgramPath)
	err := agent.Stop()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] stopping agent failed. err: %s\n", err))
	}
}
