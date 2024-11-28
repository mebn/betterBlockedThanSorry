package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/database"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
	"github.com/mebn/betterBlockedThanSorry/internal/service"
	"github.com/mebn/betterBlockedThanSorry/internal/updater"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx    context.Context
	daemon service.ServiceInterface
	db     database.DB
}

func NewApp() *App {
	daemon := service.NewBackgroundService(env.DaemonName, env.ProgramPath, service.Daemon)

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

	// move daemon to shared folder to make it
	// harder for user to uninstall application
	err := env.MoveProgram()
	if err != nil {
		fmt.Println(err)
	}

	endtime, _ := a.db.GetEndtime()
	currentTime := time.Now().Unix()
	if currentTime > endtime {
		updater, err := updater.NewUpdater()
		if err != nil {
			fmt.Println("creating updater failed:", err)
			return
		}

		wailsConfig, err := ParseWailsConfig()
		if err != nil {
			fmt.Println("json parsing failed:", err)
			return
		}

		// check version and update if needed (only in prod)
		if env.SkipUpdate {
			return
		}

		// TODO: move this to updaterAgent. pass current version somehow
		if !updater.UpToDate(wailsConfig.Info.ProductVersion) {
			updateAgent := service.NewBackgroundService(env.UpdaterAgentName, env.UpdateProgramPath, service.Agent)
			err = updateAgent.Start()
			if err != nil {
				fmt.Println("failed to start update agent: ", err)
				return
			}
			runtime.Quit(ctx)
		}
	}
}

func (a *App) shutdown(ctx context.Context) {
	fmt.Println("Shuting down")
	a.ctx = ctx
	a.db.CloseDB()
}

// daemon stuff

func (a *App) StartBlocker(blocktime int, blocklist []string) int64 {
	if blocktime == 0 {
		return 0
	}

	endtime, err := a.db.GetEndtime()
	if err != nil {
		return 0
	}

	currentTime := time.Now().Unix()

	if endtime >= currentTime {
		return 0
	}

	err = a.daemon.Start()
	if err != nil {
		return 0
	}

	endtime = time.Now().Add(time.Second * time.Duration(blocktime)).Unix()

	err = a.db.SetEndtime(endtime)
	if err != nil {
		return 0
	}

	return endtime
}

func (a *App) GetCurrentTime() int64 {
	return time.Now().Unix()
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
