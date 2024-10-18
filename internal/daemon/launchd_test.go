package daemon

import (
	"os"
	"testing"
	"time"
)

func TestCreateConfigFile(t *testing.T) {
	launchd := NewLaunchd("ignoremeiamjustasillylittletestthing", "someprogram")

	err := launchd.createConfigFile("abc", "123")
	if err != nil {
		t.Fatal("createConfigFile() failed: ", err)
	}

	got, err := os.ReadFile(launchd.path)
	if err != nil {
		t.Fatal("Read file failed: ", err)
	}

	want := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>ignoremeiamjustasillylittletestthing</string>
		<key>ProgramArguments</key>
		<array>
			<string>someprogram</string>
			<string>abc</string>
			<string>123</string>
		</array>
		<key>KeepAlive</key>
		<true/>
		<key>RunAtLoad</key>
		<true/>
	</dict>
</plist>
`

	if string(got) != want {
		t.Fatal("got:", string(got), "want:", want)
	}

	// delete file when the test is done, don't care if it fails
	os.Remove(launchd.path)
}

func TestStartAndIsRunningAndDelete(t *testing.T) {
	launchd := NewLaunchd("ignoremeiamjustasillylittletestthing2", "touch")

	// start by making sure the file dosen't exist and it's unloaded
	launchd.Stop()
	os.Remove(launchd.path)
	os.Remove("/tmp/ignoremeiamjustasillylittletestthing2")

	// start (simulate start)
	err := launchd.createConfigFile("/tmp/ignoremeiamjustasillylittletestthing2")
	if err != nil {
		t.Fatal("Create config file failed:", err)
	}

	// start the daemon
	err = launchd.startDaemon()
	if err != nil {
		t.Fatal("startDaemon() failed:", err)
	}

	// need to sleep here so daemon can create the file
	time.Sleep(time.Second)

	// does the file exist now? i.e. did the daemon work
	_, err = os.ReadFile("/tmp/ignoremeiamjustasillylittletestthing2")
	if err != nil {
		t.Fatal("File was not created, daemon didn't run:", err)
	}

	isRunning := launchd.IsRunning()

	if isRunning != STOPPED {
		t.Fatal("Daemon not running, but should be: ", err)
	}

	// finally stop the daemon
	err = launchd.Stop()
	if err != nil {
		t.Fatal("Stopping the daemon failed:", err)
	}

	// finally remove the temp test file
	os.Remove(launchd.path)
}
