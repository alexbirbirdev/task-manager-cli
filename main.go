package main

import (
	"fmt"
	"os"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   string
	updatedAt   string
}

func main() {
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
		title := os.Args[2]
		fmt.Println(title)
		//Исполнение функционала добавления

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run update <id> <title>")
			os.Exit(1)
		}
		id := os.Args[2]
		title := os.Args[3]
		fmt.Printf("Updating task with ID %s to title: %s\n", id, title)
		//Исполнение функционала изменения заголовка

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run delete <id>")
			os.Exit(1)
		}
		id := os.Args[2]
		fmt.Printf("Deleting task with ID %s\n", id)
		//Исполнение функционала удаления

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run mark-in-progress <id>")
			os.Exit(1)
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run mark-in-progress <id>")
			os.Exit(1)
		}
	case "list":
		if len(os.Args) == 2 {
			fmt.Println("list all tasks")
		}
		if len(os.Args) == 3 {
			fmt.Println("list only <status> tasks")
		}
	default:
		fmt.Printf("Invalid command: %s\n", command)
		os.Exit(1)
	}
}
