package daemon

import (
	"testing"
)

func TestCreateFileContent(t *testing.T) {
	launchd := NewLaunchd("ignoremeiamjustasillylittletestthing", "someprogram")

	fileContent := launchd.createFileContent(123)

	want := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>ignoremeiamjustasillylittletestthing.plist</string>
    <key>ProgramArguments</key>
    <array>
        <string>someprogram</string>
        <string>123</string>
    </array>
    <key>KeepAlive</key>
    <true/>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>`

	if want != fileContent {
		t.Fatal(fileContent, want)
	}
}

// func TestStartAndDeleteFile(t *testing.T) {
// 	launchd := NewLaunchd("ignoremeiamjustasillylittletestthing", "someprogram")

// }
