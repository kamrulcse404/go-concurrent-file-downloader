package commands

import (
	"concurrentfiledownloader/services"
	"fmt"
)

func HandleDownload(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("missing download URL")
	}

	downloadUrls := args[2:]

	for _, downloadUrl := range downloadUrls {
		if err := services.DownloadFile(downloadUrl); err != nil {
			return err
		}
	}

	return nil
}
