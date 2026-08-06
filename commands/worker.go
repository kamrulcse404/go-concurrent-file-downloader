package commands

import (
	"concurrentfiledownloader/services"
	"sync"
)

func worker(wg *sync.WaitGroup, jobs <-chan string, ch chan<- error) {
	defer wg.Done()

	for downloadURL := range jobs {
		err := services.DownloadFile(downloadURL)
		if err != nil {
			ch <- err
		}
	}
}
