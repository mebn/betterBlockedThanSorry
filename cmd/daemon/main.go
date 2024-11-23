package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/blocker"
	"github.com/mebn/betterBlockedThanSorry/internal/database"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

func main() {
	// get data from database
	db, err := database.NewDB(env.DBPath)
	os.WriteFile("/tmp/bbtslogger", []byte(fmt.Sprintf("db error: %s\n", err)), 0644)
	handleError(err, "Failed to load the database")
	defer db.CloseDB()

	endtime, err := db.GetEndtime()
	handleError(err, "Failed to get endtime")

	blocklist, err := db.GetBlocklist()
	handleError(err, "Failed to get blocklist")

	os.WriteFile("/tmp/bbtslogger", []byte(fmt.Sprintln(endtime, blocklist)), 0644)

	// prepare /etc/hosts file
	etcHosts := blocker.NewEtcHosts(env.EtcHostsPath, blocklist)
	etcHosts.AddBlock()

	// Timer for the overall duration
	duration := time.Until(time.Unix(endtime, 0))
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
			// TODO: write endtime back to database
			// incase user manually edits while blocker
			// is running?

			if etcHosts.IsTamperedWith() {
				etcHosts.RemoveBlock()
				etcHosts.AddBlock()
			}
		}
	}
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Printf("[ERR] %s: %s", msg, err)
		return
	}
}
