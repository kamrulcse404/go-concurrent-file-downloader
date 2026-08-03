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
		go func(downloadURL string) {
			defer wg.Done()
			if err := services.DownloadFile(downloadURL); err != nil {
				ch <- err
			}
		}(downloadURL)
	}

	wg.Wait()
	close(ch)

	var hasError bool
	for err := range ch {
		if err != nil {
			fmt.Println("Error:", err)
			hasError = true
		}
	}

	if hasError {
		return fmt.Errorf("one or more downloads failed")
	}

	return nil
}
