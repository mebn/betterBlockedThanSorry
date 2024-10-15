package main

import (
	"fmt"
	"os"

	"github.com/mebn/betterBlockedThanSorry/internal/daemon"
)

func main() {
	initSystem, err := daemon.GetDeamon()
	if err != nil {
		fmt.Println("ERR:", err)
		os.Exit(1)
	}

	isRunning, err := initSystem.IsRunning()
}
