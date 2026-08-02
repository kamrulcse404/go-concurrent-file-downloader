package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

const (
	dataDir = "downloads"
)

func DownloadFile(rawUrl string) error {
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

	_, err = os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
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

	fmt.Println("Downloaded:", fileName)

	return nil
}
