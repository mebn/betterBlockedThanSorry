package initsystem

import (
	"errors"
	"fmt"
	"runtime"
)

type InitSystemType interface {
	Start(args ...string) error
	Stop() error
	IsRunning() bool
}

func NewDaemon(daemonName, programName string) (InitSystemType, error) {
	os := runtime.GOOS

	switch os {
	case "darwin":
		return newLaunchd(daemonName, programName), nil
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
	}

	return nil, errors.New("OS not supported: " + os)
}
