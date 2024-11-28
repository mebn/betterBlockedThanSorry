package updater

import (
	"encoding/json"
	"fmt"
)

type gitHubRelease struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func newGitHubRelease() (gitHubRelease, error) {
	url := "https://api.github.com/repos/mebn/betterBlockedThanSorry/releases/latest"

	body, err := getRequest(url, map[string]string{
		"Authorization": "token " + apiKey,
	})
	if err != nil {
		return gitHubRelease{}, fmt.Errorf("get request failed: %v", err)
	}
	defer body.Close()

	var release gitHubRelease
	err = json.NewDecoder(body).Decode(&release)
	if err != nil {
		return gitHubRelease{}, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return release, nil
}
