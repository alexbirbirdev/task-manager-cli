package commands

import (
	"fmt"
	"task-manager-cli/config"
	"task-manager-cli/models"
	"task-manager-cli/utils"
	"time"
)

func MarkDone(tasks []models.Task, id int) error {
	var found bool
	for i, task := range tasks {
		if task.ID == id {
			found = true
			task.Status = "Done"
			task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			tasks[i] = task
		}
	}

	if !found {
		fmt.Println("Task not found")
	}

	// Сохраняем изменения в JSON-файл
	err := utils.SaveFile(config.TasksFilePath, tasks)
	if err != nil {
		return err
	}

	fmt.Println("Status was updated")
	return nil
}
