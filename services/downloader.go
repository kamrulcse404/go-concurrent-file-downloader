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

func DownloadFile(rawUrl string, wg *sync.WaitGroup, ch chan error) {
	defer wg.Done()

	resp, err := http.Get(rawUrl)

	if err != nil {
		ch <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- fmt.Errorf("download failed: %s", resp.Status)
		return
	}

	u, err := url.Parse(rawUrl)

	if err != nil {
		ch <- fmt.Errorf("parse url: %w", err)
		return
	}

	fileName := path.Base(u.Path)
	if fileName == "/" || fileName == "." {
		fileName = "downloaded_file"
	}

	if err = ensureDownloadDir(dataDir); err != nil {
		ch <- err
		return
	}

	filePath := filepath.Join(dataDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		ch <- err
		return
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		ch <- err
		return
	}

	fmt.Println("Downloaded:", filePath)
}
