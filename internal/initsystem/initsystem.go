package initsystem

import (
	"runtime"
)

type InitSystemType interface {
	Start(args ...string) error
	Stop() error
	IsRunning() bool
}

func NewDaemon(daemonName, programName string) InitSystemType {
	os := runtime.GOOS

	if os == "darwin" {
		return newLaunchd(daemonName, programName)
	}

	// windows
	return nil
}
