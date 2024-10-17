package daemon

import (
	"fmt"
	"os"
	"os/exec"
)

type Launchd struct {
	program    string
	daemonName string
	pathGlobal string
	pathUser   string
}

func NewLaunchd(daemonName, program string) *Launchd {
	return &Launchd{
		program:    program,
		daemonName: daemonName,
		pathGlobal: "/Library/LaunchDaemons/",
		pathUser:   "~/Library/LaunchAgents/",
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
	// create file
	fileName := fmt.Sprintf("%s.plist", l.daemonName)
	fullPath := fmt.Sprintf("%s%s", l.pathGlobal, fileName)

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
	fmt.Println(fileName)
	err = exec.Command("launchctl", "load", fileName).Run() // this wont load!?
	if err != nil {
		return err
	}

	return nil
}

func (l *Launchd) IsRunning() (bool, error) {
	return true, nil
}

func (l *Launchd) DeleteFile() error {
	path := fmt.Sprintf("%s%s", l.pathGlobal, l.daemonName)

	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
