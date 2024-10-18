package main

import (
	"fmt"
	"os"

	"github.com/mebn/betterBlockedThanSorry/internal/daemon"
)

func main() {
	initSystem, err := daemon.NewDeamon(
		"com.bbts.daemon",
		"/Users/mebn/go/bin/bbts_daemon",
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// newTime := time.Now().Add(time.Second * 5).Unix()

	// err = initSystem.Start(newTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	isRunning := initSystem.IsRunning()
	fmt.Println(isRunning)
}
