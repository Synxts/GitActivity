package main

import (
	"activity/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: <usarname>")
		return
	}

	username := strings.Join(args[1:], " ")

	events, err := utils.FetchGitHubActivity(username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	utils.DisplayActicityEvent(events)
}
