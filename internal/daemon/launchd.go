package daemon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

type daemonOrAgentType int

const (
	daemon daemonOrAgentType = iota
	agent
)

type Launchd struct {
	programToExecute    string
	nameOfDaemonOrAgent string
	pathToDaemonOrAgent string
	daemonOrAgent       daemonOrAgentType
}

func newLaunchdDaemon(nameOfDaemon, program string) *Launchd {
	return &Launchd{
		programToExecute:    program,
		nameOfDaemonOrAgent: nameOfDaemon,
		pathToDaemonOrAgent: strings.ReplaceAll(
			fmt.Sprintf("/Library/LaunchDaemons/%s.plist", nameOfDaemon),
			`"`, `\"`),
		daemonOrAgent: daemon,
	}
}

func newLaunchdAgent(nameOfAgent, program string) *Launchd {
	return &Launchd{
		programToExecute:    program,
		nameOfDaemonOrAgent: nameOfAgent,
		pathToDaemonOrAgent: strings.ReplaceAll(
			fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", env.Home(), nameOfAgent),
			`"`, `\"`),
		daemonOrAgent: agent,
	}
}

func (l *Launchd) Start(args ...string) error {
	fileContent, err := l.createConfigFile(args...)
	if err != nil {
		return err
	}

	formattedFileContent := strings.ReplaceAll(fileContent, `"`, `\"`)

	// agent
	script := fmt.Sprintf(
		`echo '%s' > \"%s\" && launchctl load -w \"%s\"`,
		formattedFileContent, l.pathToDaemonOrAgent, l.pathToDaemonOrAgent)
	cmd := exec.Command("bash", "-c", script)

	// daemon
	if l.daemonOrAgent == daemon {
		script = fmt.Sprintf(
			`do shell script "%s" with administrator privileges`,
			script)
		cmd = exec.Command("osascript", "-e", script)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error in Start(). out: %s, err: %s\n", out, err)
		return err
	}

	return nil
}

func (l *Launchd) Stop() error {
	// agent
	script := fmt.Sprintf(`launchctl unload -w \"%s\" && rm -fr \"%s\"`, l.pathToDaemonOrAgent, l.pathToDaemonOrAgent)
	cmd := exec.Command("bash", "-c", script)

	// daemon
	if l.daemonOrAgent == daemon {
		script = fmt.Sprintf(
			`do shell script "%s" with administrator privileges`,
			script)
		cmd = exec.Command("osascript", "-e", script)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error in Stop(). out: %s, err: %s\n", out, err)
		return err
	}

	return nil
}

func (l *Launchd) createConfigFile(args ...string) (string, error) {
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
`, l.nameOfDaemonOrAgent, l.programToExecute, formattedArgs)

	return fileContent, nil
}
