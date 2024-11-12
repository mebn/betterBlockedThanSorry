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

	// some setup
	writeInfoToFile("/tmp/bbts.log", fmt.Sprintf("%d", endTime))
	// etcHosts.RemoveBlock()
	etcHosts.AddBlock()

	// Timer for the overall duration
	duration := time.Until(time.Unix(endTime, 0))
	durationTimer := time.NewTimer(duration)
	defer durationTimer.Stop()

	// Ticker for the periodic checks
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-durationTimer.C:
			etcHosts.RemoveBlock()
			return
		case <-ticker.C:
			writeInfoToFile("/tmp/bbts.log", fmt.Sprintf("%d", endTime))

			if etcHosts.IsTamperedWith() {
				etcHosts.RemoveBlock()
				etcHosts.AddBlock()
			}
		}
	}
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

func writeInfoToFile(filename, content string) error {
	currentData, _ := os.ReadFile(filename)

	if string(currentData) != content {
		_ = os.WriteFile(filename, []byte(content), 0644)
	}

	return nil
}
