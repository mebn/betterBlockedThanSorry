package initsystem

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Launchd struct {
	programName string
	daemonName  string
	daemonPath  string
}

func NewLaunchd(daemonName, program string) *Launchd {
	return &Launchd{
		programName: program,
		daemonName:  daemonName,
		daemonPath:  fmt.Sprintf("/Library/LaunchDaemons/%s.plist", daemonName),
	}
}

func (l *Launchd) Start(args []string) error {
	err := l.createConfigFile(args)
	if err != nil {
		return err
	}

	err = exec.Command("launchctl", "load", "-w", l.daemonPath).Run()
	if err != nil {
		return err
	}

	return nil
}

func (l *Launchd) Stop() error {
	err := exec.Command("launchctl", "unload", "-w", l.daemonPath).Run()
	if err != nil {
		return err
	}

	err = l.deleteFile()
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

func (l *Launchd) createConfigFile(args []string) error {
	file, err := os.OpenFile(l.daemonPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// handle args
	formattedArgs := ""
	for _, arg := range args {
		formattedArgs = fmt.Sprintf("%s\n\t\t\t<string>%s</string>", formattedArgs, arg)
	}

	fileContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>%s</string>
		<key>ProgramArguments</key>
		<array>
			<string>%s</string>%s
		</array>
		<key>KeepAlive</key>
		<true/>
		<key>RunAtLoad</key>
		<true/>
	</dict>
</plist>
`, l.daemonName, l.programName, formattedArgs)

	// write to file
	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	return nil
}

func (l *Launchd) deleteFile() error {
	err := os.Remove(l.daemonPath)
	if err != nil {
		return err
	}

	return nil
}
