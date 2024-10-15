package daemon

const launchdPlistTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>{{.Label}}</string>
    <key>ProgramArguments</key>
    <array>
        <string>{{.ExecStart}}</string>
    </array>
    <key>KeepAlive</key>
    <true/>
    <key>RunAtLoad</key>
    <true/>
    <key>StandardErrorPath</key>
    <string>{{.ErrorLog}}</string>
    <key>StandardOutPath</key>
    <string>{{.OutputLog}}</string>
</dict>
</plist>`

type Launchd struct {
	path    string
	content string
}

func NewLaunchd() *Launchd {
	return &Launchd{
		path:    "",
		content: "",
	}
}

func (l *Launchd) Start(timeOffset int) error {
	return nil
}

func (l *Launchd) IsRunning() (bool, error) {
	return true, nil
}

func (l *Launchd) DeleteFile() error {
	return nil
}
