package main

import (
	"fmt"
	"os"
	"strconv"
	"task-manager-cli/commands"
	"task-manager-cli/config"
	"task-manager-cli/models"
	"task-manager-cli/utils"
)

func main() {
	var tasks []models.Task

	if err := utils.LoadTasks(config.TasksFilePath, &tasks); err != nil {
		fmt.Println("Error loading file with tasks", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [string]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run add <title>")
			os.Exit(1)
		}
		description := os.Args[2]
		commands.AddTask(description, tasks)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run update <id> <title>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		description := os.Args[3]
		commands.UpdateTask(tasks, id, description)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run delete <id>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		commands.DeleteTask(tasks, id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run mark-in-progress <id>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		commands.MarkInProgress(tasks, id)

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run mark-in-progress <id>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		commands.MarkDone(tasks, id)

	case "list":
		var status string
		if len(os.Args) == 3 {
			status = os.Args[2]
		}
		commands.ListTasks(tasks, status)

	default:
		fmt.Printf("Invalid command: %s\n", command)
		os.Exit(1)
	}
}
