package blocker

import (
	"fmt"
	"os"
)

func OpenLogFile() *os.File {
	logFile, err := os.OpenFile("/tmp/bbts_daemon.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}

	return logFile
}

func OpenFile(filename string, logFile *os.File) *os.File {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		if logFile != nil {
			logFile.WriteString(fmt.Sprintf("[ERR] Can't open %s file. \n", filename))
		}
		os.Exit(1)
	}

	return file
}