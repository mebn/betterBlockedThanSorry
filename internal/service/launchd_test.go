//go:build darwin

package service

import (
	"os"
	"path/filepath"
	"testing"
	"time"
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

func TestStartAndStop(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("failed to create temp dir")
	}
	defer os.RemoveAll(tempDir)
	fileToCreate := filepath.Join(tempDir, "a")

	a := newLaunchd("com.betterblockedthan.sorry.test457", "touch", Agent)

	t.Run("touch a file (with args)", func(t *testing.T) {
		a.Start(fileToCreate)
		time.Sleep(500 * time.Millisecond)
		a.Stop()

		_, err = os.Stat(fileToCreate)
		if err != nil {
			t.Fatal("failed to create file, start() failed")
		}

		// is plist gone?
		_, err = os.Stat(a.pathToDaemonOrAgent)
		if err == nil {
			t.Fatal("the plist file should be gone, but isn't")
		}
	})
}
