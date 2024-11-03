package initsystem

import (
	"errors"
	"fmt"
	"runtime"
)

type RunningStatus string

const (
	RUNNING     RunningStatus = "running"
	NOT_RUNNING RunningStatus = "not running"
	STOPPED     RunningStatus = "stopped"
)

type InitSystemType interface {
	Start(args []string) error
	Stop() error
	IsRunning() RunningStatus
}

func NewDaemon(daemonName, programName string) (InitSystemType, error) {
	os := runtime.GOOS

	switch os {
	case "darwin":
		return NewLaunchd(daemonName, programName), nil
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
	}

	return nil, errors.New("OS not supported: " + os)
}
