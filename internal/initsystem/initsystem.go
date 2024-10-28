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
	Start(newTime int64) error
	Stop() error
	IsRunning() RunningStatus
	DeleteFile() error
}

func NewDaemon(daemonName, program string) (InitSystemType, error) {
	os := runtime.GOOS

	switch os {
	case "darwin":
		return NewLaunchd(daemonName, program), nil
	case "linux":
		// return Systemd{}, nil
	case "windows":
		fmt.Println("windows")
	}

	return nil, errors.New("OS not supported: " + os)
}
