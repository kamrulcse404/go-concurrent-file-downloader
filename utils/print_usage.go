package utils

import "fmt"

func PrintUsage() {
	fmt.Println()
	fmt.Println("Concurrent File Downloader")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  downloader download <url>")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  downloader download https://example.com/file.zip")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help      Show this help message")
	fmt.Println()
}
