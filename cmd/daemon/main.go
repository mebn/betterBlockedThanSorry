package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/blocker"
	"github.com/mebn/betterBlockedThanSorry/internal/database"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
	"github.com/mebn/betterBlockedThanSorry/internal/service"
)

func main() {
	file, _ := os.OpenFile("/tmp/bbtsblocker.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	file.WriteString("deamon starting\n")

	// get data from database
	db, err := database.NewDB(env.DBPath)
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to load the database: %s\n", err))
		stop(file)
	}
	defer db.CloseDB()

	// sleep so we have time to write and then read endtime
	time.Sleep(1 * time.Second)

	endtime, err := db.GetEndtime()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to get endtime: %s\n", err))
		stop(file)
	}

	blocklist, err := db.GetBlocklist()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] Failed to get blocklist: %s\n", err))
		stop(file)
	}

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
			stop(file)
		case <-ticker.C:
			currentTime := time.Now().Unix()
			if currentTime >= endtime {
				etcHosts.RemoveBlock()
				stop(file)
			}
			// TODO: handle new urls addad after start
			if etcHosts.IsTamperedWith() {
				etcHosts.RemoveBlock()
				etcHosts.AddBlock()
			}
		}
	}
}

func stop(file *os.File) {
	daemon := service.NewBackgroundService(env.DaemonName, env.ProgramPath, service.Daemon)
	err := daemon.Stop()
	if err != nil {
		file.WriteString(fmt.Sprintf("[ERR] stopping daemon failed. err: %s\n", err))
	}
}
