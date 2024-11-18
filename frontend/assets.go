// This file is used to get the assets to the cmd/main/ package.

package frontend

import "embed"

//go:embed all:dist
var Assets embed.FS
