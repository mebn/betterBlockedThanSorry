package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed wails.json
var wailsJSON []byte

// WailsConfig represents the structure of the wails.json file.
type WailsConfig struct {
	Schema          string `json:"$schema"`
	Name            string `json:"name"`
	BuildDir        string `json:"build:dir"`
	FrontendDir     string `json:"frontend:dir"`
	FrontendInstall string `json:"frontend:install"`
	FrontendBuild   string `json:"frontend:build"`
	FrontendDev     struct {
		Watcher   string `json:"frontend:dev:watcher"`
		ServerURL string `json:"frontend:dev:serverUrl"`
	} `json:"-"`
	Author struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
	Info struct {
		ProductVersion string `json:"productVersion"`
		Copyright      string `json:"copyright"`
		Comments       string `json:"comments"`
	} `json:"info"`
}

func ParseWailsConfig() (*WailsConfig, error) {
	var config WailsConfig
	if err := json.Unmarshal(wailsJSON, &config); err != nil {
		return nil, fmt.Errorf("failed to parse wails.json: %w", err)
	}
	return &config, nil
}
