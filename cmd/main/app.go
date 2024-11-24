package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/database"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
	"github.com/mebn/betterBlockedThanSorry/internal/initsystem"
)

type App struct {
	ctx    context.Context
	daemon initsystem.InitSystemType
	db     database.DB
}

func NewApp() *App {
	// TODO: check if daemon is already running?

	daemon, err := initsystem.NewDaemon(env.DaemonName, env.ProgramPath)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %s", err))
	}

	db, err := database.NewDB(env.DBPath)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %s", err))
	}

	return &App{
		daemon: daemon,
		db:     db,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	a.ctx = ctx
	a.db.CloseDB()
}

// daemon stuff

func (a *App) StartBlocker(blocktime int, blocklist []string) int64 {
	if blocktime == 0 {
		return 0
	}

	isRunning := a.daemon.IsRunning()
	if isRunning {
		return 0
	}

	endtime := time.Now().Add(time.Second * time.Duration(blocktime)).Unix()

	err := a.db.SetEndtime(endtime)
	if err != nil {
		return 0
	}

	err = a.daemon.Start()
	if err != nil {
		return 0
	}

	return endtime
}

func (a *App) GetDaemonRunningStatus() bool {
	isRunning := a.daemon.IsRunning()

	if !isRunning {
		_ = a.daemon.Stop()
	}

	return isRunning
}

// endtime stuff

func (a *App) GetEndtimeDB() int64 {
	endtime, err := a.db.GetEndtime()
	if err != nil {
		return 0
	}

	return endtime
}

// blocklist stuff

func (a *App) GetBlocklistDB() []string {
	blocklist, err := a.db.GetBlocklist()
	if err != nil {
		return []string{}
	}

	return blocklist
}

func (a *App) SetBlocklistDB(blocklist []string) {
	// TOOD: error handling
	err := a.db.SetBlocklist(blocklist)
	if err != nil {
	}
}
