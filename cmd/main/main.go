package main

import (
	"github.com/mebn/betterBlockedThanSorry/frontend"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "BetterBlockedThanSorry",
		Width:         650,
		Height:        510,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: frontend.Assets,
		},
		BackgroundColour: &options.RGBA{R: 242, G: 242, B: 242, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
