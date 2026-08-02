package main

import (
	"concurrentfiledownloader/commands"
	"concurrentfiledownloader/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		utils.PrintUsage()
		return
	}
	command := strings.ToLower(os.Args[1])

	switch command {
	case "-h", "--help":
		utils.PrintUsage()

	case "download":
		if err := commands.HandleDownload(os.Args); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			utils.PrintUsage()
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", command)
		utils.PrintUsage()
		os.Exit(1)
	}
}
