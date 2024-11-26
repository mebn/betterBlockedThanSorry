package main

import (
	"fmt"

	"github.com/mebn/betterBlockedThanSorry/frontend"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	// check version and upgrade if needed

	// move daemon to shared folder to make it
	// harder for user to uninstall application
	err := env.MoveProgram()
	if err != nil {
		fmt.Println(err)
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:         "BetterBlockedThanSorry",
		Width:         650,
		Height:        510,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: frontend.Assets,
		},
		BackgroundColour: &options.RGBA{R: 242, G: 242, B: 242, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
