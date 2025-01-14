package commands

import (
	"fmt"
	"strings"
	"task-manager-cli/models"
)

func ListTasks(tasks []models.Task, status string) {
	// Если никаких задач нет
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}
	c := 0

	// Выводим список задач с указанным статусом или все вместе
	for _, task := range tasks {
		if status == "" || task.Status == strings.ToLower(status) {
			fmt.Printf("ID: %v\n Description: %s\n Status: %s\n CreatedAt: %s\n UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
			c++
		}
	}

	// Если с указанным статусом нет задач
	if c == 0 {
		fmt.Println("No tasks with status", status)
	}
}
