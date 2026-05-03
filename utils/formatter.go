package utils

import (
	"activity/models"
	"fmt"
	"strings"
)

func DisplayActicityEvent(events []models.Event) {
	if len(events) == 0 {
		fmt.Println("no recent activity found.")
		return
	}

	for _, event := range events {
		var action string

		switch event.Type {
		case "PushEvent":
			commitCount := len(event.Payload.Commits)
			action = fmt.Sprintf("Pushed %d commits to %s", commitCount, event.Repo.Name)
		case "IssueEvent":
			action = fmt.Sprintf("an issue in %s", event.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("Started %s", event.Repo.Name)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", event.Repo.Name)
		case "CreateEvent":
			action = fmt.Sprintf("created %s in %s", event.Payload.RefType, event.Repo.Name)
		default:
			typeName := strings.TrimSuffix(event.Type, "Event")
			action = fmt.Sprintf("%s is %s ", typeName, event.Repo.Name)
			fmt.Println("events ends")

		}
		fmt.Printf("%s\n", action)

	}

}
