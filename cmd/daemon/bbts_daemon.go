package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const STARTBLOCK string = "# [START] generated by BBTS, do not remove or modify."
const ENDBLOCK string = "# [END] generated by BBTS, do not remove or modify."

func main() {
	// open log file
	logFile, err := os.OpenFile("/tmp/bbts_daemon.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}
	defer logFile.Close()

	// open /etc/hosts file
	// etcHostsFile, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	etcHostsFile, err := os.OpenFile("/tmp/etc", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logFile.WriteString("[ERR] Can't open /etc/hosts file. \n")
		os.Exit(1)
	}
	defer etcHostsFile.Close()

	// handle args
	if len(os.Args) < 2 {
		logFile.WriteString("[ERR] No argument(s) received. \n")
		os.Exit(1)
	}

	endTime, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		logFile.WriteString("[ERR] First argument is not a int. \n")
		os.Exit(1)
	}

	currentTime := time.Now().Unix()
	blocklist := []string {"reddit.com"} // this should be a args
	blockPart := generateEtcHosts(blocklist)

	// start of program/daemon
	removeBlock(etcHostsFile)
	addBlock(etcHostsFile, blockPart)

	for currentTime < endTime {
		if fileTamperedWith(etcHostsFile, blockPart) {
			removeBlock(etcHostsFile)
			addBlock(etcHostsFile, blockPart)
		}

		time.Sleep(time.Second * 1)
		currentTime = time.Now().Unix()
	}

	removeBlock(etcHostsFile)
}

func generateEtcHosts(blocklist []string) string {
	var etcHostsPart strings.Builder

	etcHostsPart.WriteString(STARTBLOCK)
	etcHostsPart.WriteString("\n")
	
	for _, url := range blocklist {
		etcHostsPart.WriteString(fmt.Sprintf("127.0.0.1 %s \n", url))
		etcHostsPart.WriteString(fmt.Sprintf("127.0.0.1 www.%s \n", url))
	}
	
	etcHostsPart.WriteString(ENDBLOCK)
	etcHostsPart.WriteString("\n")

	return etcHostsPart.String()
}

func fileTamperedWith(file *os.File, blockPart string) bool {
	resetReaderPointer(file)

	var foundBlockPart strings.Builder
	shouldWrite := false
	scanner := bufio.NewScanner(file)
	
    for scanner.Scan() {
        text := scanner.Text()
		if text == STARTBLOCK {
			shouldWrite = true
		}

		if shouldWrite {
			foundBlockPart.WriteString(text)
			foundBlockPart.WriteString("\n")
		}

		if text == ENDBLOCK {
			shouldWrite = false
		}
    }

	return foundBlockPart.String() != blockPart
}

func addBlock(file *os.File, content string) {
	resetReaderPointer(file)
	file.WriteString(content)
}

func removeBlock(file *os.File) {
	resetReaderPointer(file)

	var lines []string
	shouldAdd := true

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
		text := scanner.Text()

		if text == STARTBLOCK {
			shouldAdd = false
		}
		
		if shouldAdd {
			lines = append(lines, text)
		}
		
		if text == ENDBLOCK {
			shouldAdd = true
		}
    }

	file.Truncate(0)
	resetReaderPointer(file)

	for _, line := range lines {
		file.WriteString(line)
		file.WriteString("\n")
	}
}

func resetReaderPointer(file *os.File) {
	_, err := file.Seek(0, 0)
	if err != nil {
		os.Exit(1)
	}
}