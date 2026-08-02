package commands

import (
	"concurrentfiledownloader/services"
	"fmt"
)

func HandleDownload(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("missing download URL")
	}

	downloadURL := args[2]

	return services.DownloadFile(downloadURL)
}
