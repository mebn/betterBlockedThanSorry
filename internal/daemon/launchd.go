package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Launchd struct {
	program    string
	daemonName string
	path       string
}

func NewLaunchd(daemonName, program string) *Launchd {
	return &Launchd{
		program:    program,
		daemonName: daemonName,
		path:       "/Library/LaunchDaemons/",
	}
}

func (l *Launchd) createFileContent(timeOffset int64) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>

	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
	<plist version="1.0">
	<dict>

		<key>Label</key>
		<string>%s</string>
		<key>ProgramArguments</key>
		<array>
		    <string>%s</string>
		    <string>%d</string>
		</array>
		<key>KeepAlive</key>
		<true/>
		<key>RunAtLoad</key>
		<true/>

	</dict>
	</plist>
	`, l.daemonName, l.program, timeOffset)
}

func (l *Launchd) Start(newTime int64) error {
	fullPath := fmt.Sprintf("%s%s.plist", l.path, l.daemonName)

	// create file
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write to file
	fileContent := l.createFileContent(newTime)
	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	// start the daemon
	err = exec.Command("launchctl", "load", "-w", fullPath).Run()
	if err != nil {
		return err
	}

	return nil
}

func (l *Launchd) Stop() error {
	fullPath := fmt.Sprintf("%s%s.plist", l.path, l.daemonName)

	err := exec.Command("launchctl", "unload", "-w", fullPath).Run()
	if err != nil {
		return err
	}

	return nil
}

func (l *Launchd) IsRunning() RunningStatus {
	cmd := fmt.Sprintf("launchctl list | grep %s", l.daemonName)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return NOT_RUNNING
	}

	parts := strings.Fields(string(out))

	if len(parts) != 3 {
		return NOT_RUNNING
	}

	if parts[0] == "-" {
		return STOPPED
	}

	return RUNNING
}

func (l *Launchd) DeleteFile() error {
	path := fmt.Sprintf("%s%s.plist", l.path, l.daemonName)

	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
