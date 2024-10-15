package daemon

import (
	"errors"
	"fmt"
	"runtime"
)

type InitSystemType interface {
	Start(timeOffset int) error
	IsRunning() (bool, error)
	DeleteFile() error
}

func GetDeamon() (InitSystemType, error) {
	os := runtime.GOOS

	switch os {
	case "darwin":
		return NewLaunchd(), nil
	case "linux":
		// return Systemd{}, nil
	case "windows":
		fmt.Println("windows")
	}

	return nil, errors.New("OS not supported: " + os)
}
