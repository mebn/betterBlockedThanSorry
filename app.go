package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/initsystem"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) StartBlocker(blocktime int, blocklist []string) string {
	daemon, err := initsystem.NewDaemon("bbts_daemon_1337", "/Users/mebn/go/bin/bbts_daemon")

	if err != nil {
		return fmt.Sprintf("ERROR: %s", err)
	}

	endTime := time.Now().Add(time.Second * time.Duration(blocktime)).Unix()
	endTimeStr := strconv.FormatInt(endTime, 10)
	blocklist = append([]string{endTimeStr}, blocklist...)

	daemon.Start(blocklist)

	return fmt.Sprintf("Blocker started. Running for %d seconds! blocklist: %s", blocktime, blocklist)
}
