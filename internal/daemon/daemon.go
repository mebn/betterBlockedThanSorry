package daemon

import (
	"errors"
	"fmt"
	"runtime"
)

type InitSystemType interface {
	Start(newTime int64) error
	IsRunning() (bool, error)
	DeleteFile() error
}

func NewDeamon(daemonName, program string) (InitSystemType, error) {
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
