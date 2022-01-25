package update

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/kardianos/osext"
)

type (
	// ReleaseInfo represents github latest release info
	ReleaseInfo struct {
		TagName     string    `json:"tag_name"`
		Name        string    `json:"name"`
		Body        string    `json:"body"`
		Draft       bool      `json:"draft"`
		Prerelease  bool      `json:"prerelease"`
		CreatedAt   time.Time `json:"created_at"`
		PublishedAt time.Time `json:"published_at"`
		Assets      []Asset   `json:"assets"`
	}
	Asset struct {
		Name               string `json:"name"`
		Size               int    `json:"size"`
		DownloadCount      int    `json:"download_count"`
		BrowserDownloadURL string `json:"browser_download_url"`
	}
)

func (r *ReleaseInfo) getDownloadURL(name string) string {
	for _, a := range r.Assets {
		if a.Name == name {
			return a.BrowserDownloadURL
		}
	}
	return ""
}

// fetchReleaseInfo return githublatest release info
func fetchReleaseInfo(ctx context.Context) (*ReleaseInfo, error) {
	url := "https://api.github.com/repos/thedevsaddam/docgen/releases/latest"
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("update: failed to fetch latest release")
	}

	rf := ReleaseInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&rf); err != nil {
		return nil, err
	}
	return &rf, nil
}

// updateBinary download the binary in current binary and replace the old one
func updateBinary(ctx context.Context, url string) error {
	if url == "" {
		return errors.New("update: empty download url")
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("update: failed to fetch binary file")
	}

	dir, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir, filepath.Base(os.Args[0]))
	tmpFile := fileName + ".tmp"

	if err := os.Chmod(fileName, 0777); err != nil {
		return err
	}

	if err := os.Rename(fileName, tmpFile); err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err := os.Chmod(fileName, 0777); err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		if err := os.Rename(tmpFile, fileName); err != nil {
			return err
		}
		return err
	}

	return os.Remove(tmpFile)
}
