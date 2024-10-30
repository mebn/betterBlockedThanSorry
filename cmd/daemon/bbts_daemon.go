package main

import (
	"os"
	"strconv"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/blocker"
)

func main() {
	logFile := blocker.OpenLogFile()
	defer logFile.Close()

	etcHostsFile := blocker.OpenFile("/tmp/etc", logFile)
	defer etcHostsFile.Close()

	// handle args
	endTime, blocklist := handleArgs(logFile)
	blockPart := blocker.GenerateEtcHosts(blocklist)
	currentTime := time.Now().Unix()
	
	// start of program/daemon
	blocker.RemoveBlock(etcHostsFile)
	blocker.AddBlock(etcHostsFile, blockPart)

	for currentTime < endTime {
		if blocker.FileTamperedWith(etcHostsFile, blockPart) {
			blocker.RemoveBlock(etcHostsFile)
			blocker.AddBlock(etcHostsFile, blockPart)
		}

		time.Sleep(time.Second * 1)
		currentTime = time.Now().Unix()
	}

	blocker.RemoveBlock(etcHostsFile)
}

func handleArgs(logFile *os.File) (int64, []string) {
	if len(os.Args) < 2 {
		logFile.WriteString("[ERR] No argument(s) received. \n")
		os.Exit(1)
	}

	endTime, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		logFile.WriteString("[ERR] First argument is not a int. \n")
		os.Exit(1)
	}

	// TODO: handle blocklist from the args
	blocklist := []string {"reddit.com"}
	return endTime, blocklist
}