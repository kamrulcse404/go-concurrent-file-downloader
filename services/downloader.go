package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sync"
)

const (
	dataDir = "downloads"
)

func DownloadFile(rawUrl string, wg *sync.WaitGroup) error {
	defer wg.Done()

	resp, err := http.Get(rawUrl)

	if err != nil {
		return err
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
	if fileName == "/" || fileName == "." {
		fileName = "downloaded_file"
	}

	if err = ensureDownloadDir(dataDir); err != nil {
		return err
	}

	filePath := filepath.Join(dataDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return err
	}

	fmt.Println("Downloaded:", filePath)
	fmt.Println("Single File Downloader (v1)")

	return nil
}
