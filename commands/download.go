package commands

import (
	"concurrentfiledownloader/services"
	"fmt"
	"sync"
)

func HandleDownload(args []string) error {
	var wg sync.WaitGroup
	if len(args) < 3 {
		return fmt.Errorf("missing download URL")
	}

	downloadURLs := args[2:]
	ch := make(chan error, len(downloadURLs))

	for _, downloadURL := range downloadURLs {
		wg.Add(1)
		go services.DownloadFile(downloadURL, &wg, ch)
	}

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			return err
		}
	}

	return nil
}
