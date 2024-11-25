package daemon

import (
	"runtime"
)

type DaemonInterface interface {
	Start(args ...string) error
	Stop() error
}

// runs as root
func NewDaemon(nameOfAgent, programToExecute string) DaemonInterface {
	os := runtime.GOOS

	if os == "darwin" {
		return newLaunchdDaemon(nameOfAgent, programToExecute)
	}

	// windows
	return nil
}

// runs as user
func NewAgent(nameOfAgent, programToExecute string) DaemonInterface {
	os := runtime.GOOS

	if os == "darwin" {
		return newLaunchdAgent(nameOfAgent, programToExecute)
	}

	// windows
	return nil
}
