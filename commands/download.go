package commands

import (
	"concurrentfiledownloader/config"
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
	jobs := make(chan string)

	count := config.WorkerCount

	if len(downloadURLs) < count {
		count = len(downloadURLs)
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go worker(&wg, jobs, ch)
	}

	for _, downloadURL := range downloadURLs {
		jobs <- downloadURL
	}

	close(jobs)

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
