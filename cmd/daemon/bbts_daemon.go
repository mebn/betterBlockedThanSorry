package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/blocker"
)

func main() {
	endTime, blocklist := handleArgs()
	etcHosts := blocker.NewEtcHosts("/etc/hosts", blocklist)
	currentTime := time.Now().Unix()

	etcHosts.RemoveBlock()
	etcHosts.AddBlock()

	for currentTime < endTime {
		writeInfoToFile("/tmp/bbts.log", fmt.Sprintf("%d", endTime))

		if etcHosts.IsTamperedWith() {
			etcHosts.RemoveBlock()
			etcHosts.AddBlock()
		}

		time.Sleep(60 * time.Second)
		currentTime = time.Now().Unix()
	}

	etcHosts.RemoveBlock()
}

func handleArgs() (int64, []string) {
	if len(os.Args) < 2 {
		fmt.Println("[ERR] No argument(s) received.")
		os.Exit(1)
	}

	endTime, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("[ERR] First argument is not a int.")
		os.Exit(1)
	}

	blocklist := os.Args[2:]

	return endTime, blocklist
}

func writeInfoToFile(filename, contents string) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(contents)
}
