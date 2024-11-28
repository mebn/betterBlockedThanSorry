package service

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

type DaemonOrAgentType int

const (
	Daemon DaemonOrAgentType = iota
	Agent
)

type Launchd struct {
	programToExecute    string
	nameOfDaemonOrAgent string
	pathToDaemonOrAgent string
	daemonOrAgent       DaemonOrAgentType
}

func newLaunchd(nameOfDaemonOrAgent, programToExecute string, daemonOrAgent DaemonOrAgentType) *Launchd {
	var pathToDaemonOrAgent string

	if daemonOrAgent == Daemon {
		pathToDaemonOrAgent = fmt.Sprintf("/Library/LaunchDaemons/%s.plist", nameOfDaemonOrAgent)
	} else {
		pathToDaemonOrAgent = fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", env.Home(), nameOfDaemonOrAgent)
	}

	return &Launchd{
		programToExecute,
		nameOfDaemonOrAgent,
		pathToDaemonOrAgent,
		daemonOrAgent,
	}
}

func (l *Launchd) Start(args ...string) error {
	var fileContent string
	var err error

	if l.daemonOrAgent == Daemon {
		fileContent, err = l.createConfigFile(true, true, args...)
	} else {
		fileContent, err = l.createConfigFile(false, true, args...)
	}

	if err != nil {
		return err
	}

	formattedFileContent := strings.ReplaceAll(fileContent, `"`, `\"`)

	var script string
	var cmd *exec.Cmd

	if l.daemonOrAgent == Daemon {
		// daemon
		script = fmt.Sprintf(
			`echo '%s' > \"%s\" && launchctl load -w \"%s\"`,
			formattedFileContent, l.pathToDaemonOrAgent, l.pathToDaemonOrAgent)

		script = fmt.Sprintf(
			`do shell script "%s" with administrator privileges`,
			script)
		cmd = exec.Command("osascript", "-e", script)
	} else {
		// agent
		script = fmt.Sprintf(
			`echo '%s' > '%s' && launchctl load -w '%s'`,
			formattedFileContent, l.pathToDaemonOrAgent, l.pathToDaemonOrAgent)
		cmd = exec.Command("bash", "-c", script)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error in Start(). out: %s, err: %s\n", out, err)
		return err
	}

	return nil
}

func (l *Launchd) Stop() error {
	var script string
	var cmd *exec.Cmd

	if l.daemonOrAgent == Daemon {
		// daemon
		script = fmt.Sprintf(
			`launchctl unload -w \"%s\" && rm -fr \"%s\"`,
			l.pathToDaemonOrAgent, l.pathToDaemonOrAgent)

		script = fmt.Sprintf(
			`do shell script "%s" with administrator privileges`,
			script)
		cmd = exec.Command("osascript", "-e", script)
	} else {
		// agent
		// TODO: this does not remove
		script = fmt.Sprintf(
			`launchctl unload -w '%s' && rm -fr '%s'`,
			l.pathToDaemonOrAgent, l.pathToDaemonOrAgent)
		cmd = exec.Command("bash", "-c", script)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error in Stop(). out: %s, err: %s\n", out, err)
		return err
	}

	return nil
}

func (l *Launchd) createConfigFile(keepAlive, runAtLoad bool, args ...string) (string, error) {
	// keep alive
	keepAliveString := "<false/>"
	if keepAlive {
		keepAliveString = "<true/>"
	}

	// run at load
	runAtLoadString := "<false/>"
	if runAtLoad {
		runAtLoadString = "<true/>"
	}

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
		%s
		<key>RunAtLoad</key>
		%s
	</dict>
</plist>
`, l.nameOfDaemonOrAgent, l.programToExecute, formattedArgs, keepAliveString, runAtLoadString)

	return fileContent, nil
}
