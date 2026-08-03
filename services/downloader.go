package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	dataDir    = "downloads"
	timeFormat = "20060102_150405"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

func DownloadFile(rawUrl string) error {
	resp, err := client.Get(rawUrl)

	if err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: unexpected status %s", rawUrl, resp.Status)
	}

	u, err := url.Parse(rawUrl)

	if err != nil {
		return fmt.Errorf("parse url: %w", err)

	}

	fileName := path.Base(u.Path)

	if fileName == "." || fileName == "/" {
		fileName = "downloaded_file"
	}

	timestamp := time.Now().Format(timeFormat)
	fileName = fmt.Sprintf("%s_%s", timestamp, fileName)

	if err = ensureDownloadDir(dataDir); err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)
	}

	filePath := filepath.Join(dataDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file %s: %w", filePath, err)
	}

	defer file.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("copy response body to %s: %w", filePath, err)
	}

	fmt.Println("Downloaded:", filePath)
	return nil
}
