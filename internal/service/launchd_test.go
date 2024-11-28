//go:build darwin

package service

import (
	"testing"
)

func TestCreateConfigFile(t *testing.T) {
	launchd := newLaunchd("com.betterblockedthan.sorry.testing", "someprogram", Daemon)

	t.Run("config file without arguments", func(t *testing.T) {
		got, err := launchd.createConfigFile(true, true)
		if err != nil {
			t.Fatal("createConfigFile() failed: ", err)
		}

		want := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>com.betterblockedthan.sorry.testing</string>
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
	})

	t.Run("config file with arguments", func(t *testing.T) {
		got, err := launchd.createConfigFile(true, true, "a", "b")
		if err != nil {
			t.Fatal("createConfigFile() failed: ", err)
		}

		want := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>com.betterblockedthan.sorry.testing</string>
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
	})
}

// func TestStartAndIsRunningAndDelete(t *testing.T) {
// 	// setup
// 	launchd := newLaunchd("ignoremeiamjustasillylittletestthing2", "touch")
// 	launchd.Stop()
// 	os.Remove("/tmp/ignoremeiamjustasillylittletestthing2")

// 	// start
// 	launchd.Start("/tmp/ignoremeiamjustasillylittletestthing2")

// 	// need to sleep here so daemon can create the touch file
// 	time.Sleep(time.Second)

// 	// does the file exist now? i.e. did the daemon work
// 	_, err := os.ReadFile("/tmp/ignoremeiamjustasillylittletestthing2")
// 	if err != nil {
// 		t.Fatal("File was not created, daemon didn't run:", err)
// 	}

// 	// status
// 	want := false
// 	got := launchd.IsRunning()

// 	if got != want {
// 		t.Fatal(got, want)
// 	}

// 	// stop the daemon
// 	err = launchd.Stop()
// 	if err != nil {
// 		t.Fatal("Stopping the daemon failed:", err)
// 	}

// 	// cleanup
// 	os.Remove("/tmp/ignoremeiamjustasillylittletestthing2")
// }
