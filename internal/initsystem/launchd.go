package initsystem

import (
	"fmt"
	"os/exec"
	"strings"
)

type Launchd struct {
	programName string
	daemonName  string
	daemonPath  string
}

func newLaunchd(daemonName, program string) *Launchd {
	return &Launchd{
		programName: program,
		daemonName:  daemonName,
		daemonPath:  strings.ReplaceAll(fmt.Sprintf("/Library/LaunchDaemons/%s.plist", daemonName), `"`, `\"`),
	}
}

func (l *Launchd) Start(args ...string) error {
	fileContent, err := l.createConfigFile(args...)
	if err != nil {
		return err
	}

	formattedFileContent := strings.ReplaceAll(fileContent, `"`, `\"`)

	script := fmt.Sprintf(
		`do shell script "echo '%s' > \"%s\" && launchctl load -w \"%s\"" with administrator privileges`,
		formattedFileContent, l.daemonPath, l.daemonPath)

	out, err := exec.Command("osascript", "-e", script).CombinedOutput()
	if err != nil {
		fmt.Printf("error in Start(). out: %s, err: %s\n", out, err)
		return err
	}

	return nil
}

func (l *Launchd) Stop() error {
	script := fmt.Sprintf(
		`do shell script "launchctl unload -w \"%s\" && rm -fr \"%s\"" with administrator privileges`,
		l.daemonPath, l.daemonPath,
	)

	err := exec.Command("osascript", "-e", script).Run()
	if err != nil {
		return err
	}

	return nil
}

func (l *Launchd) IsRunning() bool {
	return false
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
`, l.daemonName, l.programName, formattedArgs)

	return fileContent, nil
}
