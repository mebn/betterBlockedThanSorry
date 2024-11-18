// This file is used to get the assets to the cmd/main/ module.

package assets

import "embed"

//go:embed all:dist
var Assets embed.FS
