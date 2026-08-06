package services

import (
	"concurrentfiledownloader/config"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"
)

var client = &http.Client{
	Timeout: config.HTTPTimeout,
}

func DownloadFile(rawUrl string) error {
	u, err := url.ParseRequestURI(rawUrl)

	if err != nil {
		return fmt.Errorf("%s parse url: %w", rawUrl, err)
	}

	resp, err := client.Get(rawUrl)

	if err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: unexpected status %s", rawUrl, resp.Status)
	}

	fileName := path.Base(u.Path)

	if fileName == "." || fileName == "/" {
		fileName = "downloaded_file"
	}

	timestamp := time.Now().Format(config.TimeFormat)
	fileName = fmt.Sprintf("%s_%s", timestamp, fileName)

	if err = ensureDownloadDir(config.DataDir); err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)
	}

	filePath := filepath.Join(config.DataDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file %s: %w", filePath, err)
	}

	defer file.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("copy response body to %s: %w", filePath, err)
	}

	fmt.Printf("Downloaded %s -> %s\n", rawUrl, filePath)
	return nil
}
