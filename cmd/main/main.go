package main

import (
	"fmt"
	"os"

	"github.com/mebn/betterBlockedThanSorry/internal/initsystem"
)

func main() {
	bbts_daemon, err := initsystem.NewDaemon(
		"com.bbts.daemon",
		"/Users/mebn/go/bin/bbts_daemon",
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// newTime := time.Now().Add(time.Second * 5).Unix()

	// err = bbts_daemon.Start(newTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	isRunning := bbts_daemon.IsRunning()
	fmt.Println(isRunning)
}
