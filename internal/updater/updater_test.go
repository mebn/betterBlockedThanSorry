package updater

import (
	"os"
	"path/filepath"
	"testing"
)

// TODO
func TestNewUpdater(t *testing.T) {
}

func TestUpToDate(t *testing.T) {
	var got, want bool

	u := Updater{
		release: gitHubRelease{
			TagName: "v1.0.0",
		},
	}

	// not up to date
	got = u.UpToDate("v1.0.1")
	want = false

	if got != want {
		t.Fatal("Should not be up to date", got, want)
	}

	// up to date
	got = u.UpToDate("v1.0.0")
	want = true

	if got != want {
		t.Fatal("Should be up to date", got, want)
	}
}

// TODO: mock instead of creating files
func TestReplaceProgram(t *testing.T) {
	u := Updater{}

	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("Failed to create temp directory:", err)
	}
	defer os.RemoveAll(tempDir)

	oldPath := filepath.Join(tempDir, "a")
	newPath := filepath.Join(tempDir, "b")

	file, err := os.Create(oldPath)
	if err != nil {
		t.Fatal("Creating the test file failed.")
	}
	defer file.Close()

	err = u.ReplaceProgram(oldPath, newPath)
	if err != nil {
		t.Fatal("ReplaceProgram failed:", err)
	}

	// does new file exist?
	_, err = os.Stat(newPath)
	if err != nil {
		t.Fatal("New file does not exist:", err)
	}

	// is old file gone?
	_, err = os.Stat(oldPath)
	if err == nil {
		t.Fatal("Old file still exist:", err)
	}

	// cleanup
	os.RemoveAll(oldPath)
	os.RemoveAll(newPath)
}

// TODO
func TestRelaunchProgram(t *testing.T) {
}

func TestGetDownloadLink(t *testing.T) {
	u := Updater{
		release: gitHubRelease{
			Assets: []Asset{
				{Name: "myprogram.app.zip", Url: "https://example.com/myprogram.app.zip"},
				{Name: "myprogram.exe.zip", Url: "https://example.com/myprogram.exe.zip"},
			},
		},
	}

	u.osName = "darwin"
	downloadLink, assetName, err := u.getDownloadLink()
	if err != nil {
		t.Fatal("getDownloadLink failed:", err)
	}

	if downloadLink != "https://example.com/myprogram.app.zip" || assetName != "myprogram.app.zip" {
		t.Fatal("Got the wrong download link and name on darwin")
	}

	u.osName = "windows"
	downloadLink, assetName, err = u.getDownloadLink()
	if err != nil {
		t.Fatal("getDownloadLink failed:", err)
	}

	if downloadLink != "https://example.com/myprogram.exe.zip" || assetName != "myprogram.exe.zip" {
		t.Fatal("Got the wrong download link and name on windows")
	}
}

func TestDownloadLatestBinary(t *testing.T) {
}
