package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
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
		return fmt.Errorf("download failed: %s", err)
	}

	fileName := path.Base(u.Path)
	if fileName == "/" || fileName == "." {
		fileName = "downloaded_file"
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Downloaded: %s", fileName)

	return nil
}
