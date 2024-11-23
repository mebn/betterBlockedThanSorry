package initsystem

import (
	"os"
	"testing"
	"time"
)

func TestCreateConfigFile(t *testing.T) {
	launchd := newLaunchd("ignoremeiamjustasillylittletestthing", "someprogram")

	// test config file without arguments
	err := launchd.createConfigFile()
	if err != nil {
		t.Fatal("createConfigFile() failed: ", err)
	}

	gotB, err := os.ReadFile(launchd.daemonPath)
	if err != nil {
		t.Fatal("Read file failed: ", err)
	}
	got := string(gotB)

	want := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>ignoremeiamjustasillylittletestthing</string>
		<key>ProgramArguments</key>
		<array>
			<string>someprogram</string>
		</array>
		<key>KeepAlive</key>
		<true/>
		<key>RunAtLoad</key>
		<true/>
	</dict>
</plist>
`

	if got != want {
		t.Fatal(got, want)
	}

	// test config file with arguments
	err = launchd.createConfigFile("a", "b")
	if err != nil {
		t.Fatal("createConfigFile() failed: ", err)
	}

	gotB, err = os.ReadFile(launchd.daemonPath)
	if err != nil {
		t.Fatal("Read file failed: ", err)
	}
	got = string(gotB)

	want = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>ignoremeiamjustasillylittletestthing</string>
		<key>ProgramArguments</key>
		<array>
			<string>someprogram</string>
			<string>a</string>
			<string>b</string>
		</array>
		<key>KeepAlive</key>
		<true/>
		<key>RunAtLoad</key>
		<true/>
	</dict>
</plist>
`

	if got != want {
		t.Fatal(got, want)
	}

	// cleanup
	os.Remove(launchd.daemonPath)
}

func TestStartAndIsRunningAndDelete(t *testing.T) {
	// setup
	launchd := newLaunchd("ignoremeiamjustasillylittletestthing2", "touch")
	launchd.Stop()
	os.Remove("/tmp/ignoremeiamjustasillylittletestthing2")

	// start
	launchd.Start("/tmp/ignoremeiamjustasillylittletestthing2")

	// need to sleep here so daemon can create the touch file
	time.Sleep(time.Second)

	// does the file exist now? i.e. did the daemon work
	_, err := os.ReadFile("/tmp/ignoremeiamjustasillylittletestthing2")
	if err != nil {
		t.Fatal("File was not created, daemon didn't run:", err)
	}

	// status
	want := false
	got := launchd.IsRunning()

	if got != want {
		t.Fatal(got, want)
	}

	// stop the daemon
	err = launchd.Stop()
	if err != nil {
		t.Fatal("Stopping the daemon failed:", err)
	}

	// cleanup
	os.Remove("/tmp/ignoremeiamjustasillylittletestthing2")
}
