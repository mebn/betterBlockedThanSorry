package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/blocker"
	"github.com/mebn/betterBlockedThanSorry/internal/daemon"
	"github.com/mebn/betterBlockedThanSorry/internal/database"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

func main() {
	file, _ := os.OpenFile("/tmp/bbts.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	file.WriteString("deamon starting\n")

	// get data from database
	// problem: this is located at /var/root/.bbtsDEV/db/db
	db, err := database.NewDB(env.DBPath)
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to load the database: %s\n", err))
		return
	}
	defer db.CloseDB()

	// sleep so we have time to write and then read endtime
	time.Sleep(1 * time.Second)

	endtime, err := db.GetEndtime()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to get endtime: %s\n", err))
		return
	}

	blocklist, err := db.GetBlocklist()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to get blocklist: %s\n", err))
		return
	}

	// prepare /etc/hosts file
	etcHosts := blocker.NewEtcHosts(env.EtcHostsPath, blocklist)
	etcHosts.AddBlock()

	file.WriteString(fmt.Sprintf("[INFO] DB: %s, endtime: %d, blocklist: %v, etc: %s\n", env.DBPath, endtime, blocklist, env.EtcHostsPath))

	// Timer for the overall duration
	duration := time.Until(time.Unix(endtime, 0))
	durationTimer := time.NewTimer(duration)
	defer durationTimer.Stop()

	// Ticker for the periodic checks
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		// if computer is asleep, this won't get caught
		case <-durationTimer.C:
			etcHosts.RemoveBlock()
			err := stop(&etcHosts, file)
			if err == nil {
				return
			}
		case <-ticker.C:
			currentTime := time.Now().Unix()
			if currentTime >= endtime {
				err := stop(&etcHosts, file)
				if err == nil {
					return
				}
			}
			// TODO: handle new urls addad after start
			if etcHosts.IsTamperedWith() {
				etcHosts.RemoveBlock()
				etcHosts.AddBlock()
			}
		}
	}
}

func stop(etcHosts *blocker.EtcHosts, file *os.File) error {
	etcHosts.RemoveBlock()
	daemon := daemon.NewDaemon(env.DaemonName, env.ProgramPath)
	err := daemon.Stop()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] stopping daemon failed. err: %s\n", err))
	}
	return nil
}
