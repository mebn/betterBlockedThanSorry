package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/initsystem"
)

type App struct {
	ctx    context.Context
	daemon initsystem.InitSystemType
}

func NewApp() *App {
	daemon, err := initsystem.NewDaemon("bbts_daemon_1337", "/Users/mebn/go/bin/bbts_daemon")

	if err != nil {
		panic(fmt.Sprintf("ERROR: %s", err))
	}

	return &App{
		daemon: daemon,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartBlocker(blocktime int, blocklist []string) int64 {
	endTime := time.Now().Add(time.Second * time.Duration(blocktime)).Unix()
	endTimeStr := strconv.FormatInt(endTime, 10)
	blocklist = append([]string{endTimeStr}, blocklist...)

	err := a.daemon.Start(blocklist)
	if err != nil {
		fmt.Printf("Error starting blocker: %s\n", err)
		return 0
	}

	return endTime
}

func (a *App) GetDaemonRunningStatus() bool {
	isRunning := a.daemon.IsRunning()

	if !isRunning {
		_ = a.daemon.Stop()
	}

	return isRunning
}

func (a *App) GetEndTime(filename string) int64 {
	contentB, _ := os.ReadFile("/tmp/bbts.log")
	content := string(contentB)
	endTime, _ := strconv.ParseInt(content, 10, 64)

	return endTime
}
