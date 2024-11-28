package service

import (
	"runtime"
)

type ServiceInterface interface {
	Start(args ...string) error
	Stop() error
}

func NewBackgroundService(nameOfDaemonOrAgent, programToExecute string, daemonOrAgent DaemonOrAgentType) ServiceInterface {
	switch runtime.GOOS {
	case "darwin":
		return newLaunchd(nameOfDaemonOrAgent, programToExecute, daemonOrAgent)
	case "windows":
		return nil
	}

	return nil
}
