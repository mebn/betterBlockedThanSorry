package updater

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

// TODO: create a new key
const apiKey = "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"

type Updater struct {
	release gitHubRelease
	goOS    string
	goARCH  string
}

func NewUpdater() (Updater, error) {
	release, err := newGitHubRelease()
	if err != nil {
		return Updater{}, err
	}

	return Updater{
		release: release,
		goOS:    runtime.GOOS,
		goARCH:  runtime.GOARCH,
	}, nil
}

func (u *Updater) UpToDate(currentVersion string) bool {
	return u.release.TagName == currentVersion
}

// Download the latest version to a temporary/shared folder
func (u *Updater) DownloadLatestBinary() (string, error) {
	downloadlink, assetName, err := u.getDownloadLink()
	if err != nil {
		return "", err
	}

	zipPath, err := downloadZIP(downloadlink, assetName)
	if err != nil {
		return "", nil
	}

	if strings.HasSuffix(assetName, ".zip") {
		err := unzip(zipPath, env.DownloadPath)
		if err != nil {
			return "", fmt.Errorf("failed to unzip file: %v", err)
		}
	}

	// assume file inside zip is the same, but with .zip removed
	binaryName := assetName
	binaryName = strings.ReplaceAll(binaryName, ".amd64.zip", "")
	binaryName = strings.ReplaceAll(binaryName, ".arm64.zip", "")

	return binaryName, nil
}

func (u *Updater) ReplaceProgram(oldPath, newPath string) error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("moving the file failed. err: %s", err)
	}
	return nil
}

func (u *Updater) RelaunchProgram(path string) error {
	switch u.goOS {
	case "darwin":
		if !strings.HasSuffix(path, ".app") {
			return fmt.Errorf("invalid macOS application path: %s", path)
		}

		cmd := exec.Command("open", path)

		err := cmd.Start()
		if err != nil {
			return fmt.Errorf("failed to relaunch application: %w", err)
		}

		return nil

	case "windows":
		if !strings.HasSuffix(path, ".exe") {
			return fmt.Errorf("invalid Windows executable path: %s", path)
		}

		cmd := exec.Command(path)

		err := cmd.Start()
		if err != nil {
			return fmt.Errorf("failed to relaunch application: %w", err)
		}

		return nil
	}

	return fmt.Errorf("unsupported OS: %s", u.goOS)
}

// helpers

func (u *Updater) getDownloadLink() (string, string, error) {
	downloadLink, assetName := "", ""

	for _, asset := range u.release.Assets {
		// macos
		macosAmd64 := u.goOS == "darwin" && u.goARCH == "amd64" && strings.HasSuffix(asset.Name, ".app.amd64.zip")
		macosArm64 := u.goOS == "darwin" && u.goARCH == "arm64" && strings.HasSuffix(asset.Name, ".app.arm64.zip")

		// windows
		windowsAmd64 := u.goOS == "windows" && u.goARCH == "amd64" && strings.HasSuffix(asset.Name, ".exe.amd64.zip")
		windowsArm64 := u.goOS == "windows" && u.goARCH == "arm64" && strings.HasSuffix(asset.Name, ".exe.arm64.zip")

		if macosAmd64 || macosArm64 || windowsAmd64 || windowsArm64 {
			downloadLink = asset.Url
			assetName = asset.Name
			break
		}
	}

	if downloadLink == "" || assetName == "" {
		return "", "", fmt.Errorf("no download link found")
	}

	return downloadLink, assetName, nil
}
