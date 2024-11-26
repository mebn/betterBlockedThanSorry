package updater

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

// TODO: create a new key
const apiKey = "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"

type gitHubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"assets"`
}

type Updater struct {
	release gitHubRelease
}

func NewUpdater() (Updater, error) {
	url := "https://api.github.com/repos/mebn/betterBlockedThanSorry/releases/latest"
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Updater{}, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "token "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return Updater{}, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Updater{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var release gitHubRelease
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return Updater{}, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return Updater{
		release: release,
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

	zipPath, err := u.downloadZIP(downloadlink, assetName)
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
	binaryName := strings.ReplaceAll(assetName, ".zip", "")

	return binaryName, nil
}

func (u *Updater) ReplaceProgram(oldPath, newPath string) error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("moving the file failed. err: %s", err)
	}
	return nil
}

// helpers

func (u *Updater) getDownloadLink() (string, string, error) {
	downloadlink, assetName := "", ""

	for _, asset := range u.release.Assets {
		macos := strings.Contains(asset.Name, ".app")
		win := strings.Contains(asset.Name, ".exe")

		if macos || win {
			downloadlink = asset.Url
			assetName = asset.Name
			break
		}
	}

	if downloadlink == "" || assetName == "" {
		return "", "", fmt.Errorf("no download link found")
	}

	return downloadlink, assetName, nil
}

func (u *Updater) downloadZIP(downloadlink, assetName string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", downloadlink, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Accept", "application/octet-stream")
	req.Header.Set("Authorization", "token "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	filePath := env.SafeFile(env.DownloadPath, assetName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filePath, nil
}

func unzip(src, dest string) error {
	// Open the ZIP archive
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer r.Close()

	// Iterate through each file in the archive
	for _, file := range r.File {
		filePath := filepath.Join(dest, file.Name)

		// Create directory or file
		if file.FileInfo().IsDir() {
			// Create directory
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		} else {
			// Create file
			err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directories for file: %v", err)
			}

			outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return fmt.Errorf("failed to create file: %v", err)
			}
			defer outFile.Close()

			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("failed to open zip entry: %v", err)
			}
			defer rc.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return fmt.Errorf("failed to copy zip content: %v", err)
			}
		}
	}

	return nil
}