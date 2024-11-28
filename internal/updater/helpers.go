package updater

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

func getRequest(url string, headers map[string]string) (io.ReadCloser, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func downloadZIP(downloadlink, assetName string) (string, error) {
	body, err := getRequest(downloadlink, map[string]string{
		"Accept":        "application/octet-stream",
		"Authorization": "token " + apiKey,
	})
	if err != nil {
		return "", fmt.Errorf("get request failed: %v", err)
	}
	defer body.Close()

	filePath := env.SafeFile(env.DownloadPath, assetName)

	err = writeToLocalFile(filePath, body)
	if err != nil {
		return "", fmt.Errorf("failed to write to local file. %v", err)
	}

	return filePath, nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer r.Close()

	// Iterate through each file in the archive
	for _, file := range r.File {
		if file.FileInfo().IsDir() {
			env.SafePath(dest, file.Name)
			continue
		}

		filePath := env.SafeFile(dest, file.Name)

		// read file from zip
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open zip entry: %v", err)
		}
		defer rc.Close()

		err = writeToLocalFile(filePath, rc)
		if err != nil {
			return fmt.Errorf("failed to write to local file. %v", err)
		}
	}

	return nil
}

func writeToLocalFile(filePath string, content io.ReadCloser) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}
