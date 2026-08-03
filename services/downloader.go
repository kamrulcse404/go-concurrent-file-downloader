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
	dataDir = "downloads"
)

func DownloadFile(rawUrl string) error {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(rawUrl)

	if err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", resp.Status)
	}

	u, err := url.Parse(rawUrl)

	if err != nil {
		return fmt.Errorf("parse url: %w", err)

	}

	fileName := path.Base(u.Path)
	if fileName == "." || fileName == "/" {
		fileName = "downloaded_file"
	}

	if err = ensureDownloadDir(dataDir); err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)
	}

	filePath := filepath.Join(dataDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return fmt.Errorf("%s: %w", rawUrl, err)

	}

	fmt.Println("Downloaded:", filePath)
	return nil
}
