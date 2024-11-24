package main

import (
	"fmt"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/blocker"
	"github.com/mebn/betterBlockedThanSorry/internal/database"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

func main() {
	// get data from database
	db, err := database.NewDB(env.DBPath)
	if err != nil {
		fmt.Printf("[ERR] Failed to load the database: %s", err)
		return
	}
	defer db.CloseDB()

	endtime, err := db.GetEndtime()
	if err != nil {
		fmt.Printf("[ERR] Failed to get endtime: %s", err)
		return
	}

	blocklist, err := db.GetBlocklist()
	if err != nil {
		fmt.Printf("[ERR] Failed to get blocklist: %s", err)
		return
	}

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
			// TODO: handle new urls addad after start
			if etcHosts.IsTamperedWith() {
				etcHosts.RemoveBlock()
				etcHosts.AddBlock()
			}
		}
	}
}
