package main

import (
	"fmt"

	"github.com/mebn/betterBlockedThanSorry/internal/daemon"
	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

func createConfigFile() string {
	fileContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>%s</string>
		<key>ProgramArguments</key>
		<array>
			<string>%s</string>
		</array>
		<key>KeepAlive</key>
		<true/>
		<key>RunAtLoad</key>
		<true/>
	</dict>
</plist>
`, "somefile", "someprogram")

	return fileContent
}

func main() {
	var err error

	// check version and upgrade if needed
	updateAgent := daemon.NewAgent(env.UpdaterAgentName, env.UpdateProgramPath)
	err = updateAgent.Start()
	if err != nil {
		fmt.Println("Some error. err:", err)
	}
	// fileContent := createConfigFile()

	// formattedFileContent := strings.ReplaceAll(fileContent, `"`, `\"`)

	// script := fmt.Sprintf(
	// 	`do shell script "echo '%s' > %s" with administrator privileges`,
	// 	formattedFileContent, "/tmp/somefile")

	// out, err := exec.Command("osascript", "-e", script).CombinedOutput()
	// if err != nil {
	// 	fmt.Printf("error in Start(). out: %s, err: %s\n", out, err)
	// }

	// currentUser, err := user.Current()
	// if err != nil {
	// 	log.Fatalf("Failed to get current user: %v", err)
	// }

	// println(currentUser.Uid)

	// if currentUser.Uid != "0" {
	// 	// Get the full path to the current binary
	// 	fullPath, err := os.Executable()
	// 	if err != nil {
	// 		log.Fatalf("Failed to get executable path: %v", err)
	// 	}
	// 	escapedPath := filepath.Clean(fullPath)

	// 	// Create AppleScript command to elevate privileges
	// 	script := fmt.Sprintf(`do shell script "%s" with administrator privileges`, escapedPath)
	// 	println(script)

	// 	// osascript -e 'do shell script "/Users/mebn/marcus/code/betterBlockedThanSorry/main" with administrator privileges'
	// 	// osascript -e 'do shell script "echo fgfg" with administrator privileges'

	// 	cmd := exec.Command("osascript", "-e", script)

	// 	// Run the command and capture output
	// 	output, err := cmd.CombinedOutput()
	// 	if err != nil {
	// 		log.Fatalf("Failed to run with elevated privileges: %v\nOutput: %s", err, string(output))
	// 	}

	// 	log.Println("App successfully ran with elevated privileges!")
	// 	return
	// }

	// // Now running as root
	// for {
	// 	fmt.Println("Now running as root")
	// 	time.Sleep(1 * time.Second)
	// }
}
