package commands

import (
	"fmt"
	"task-manager-cli/config"
	"task-manager-cli/models"
	"task-manager-cli/utils"
)

func DeleteTask(tasks []models.Task, id int) error {
	var found bool
	newTasks := []models.Task{}
	for _, task := range tasks {
		if task.ID == id {
			found = true
		} else {
			newTasks = append(newTasks, task)
		}
	}

	if !found {
		fmt.Println("Task not found")
	}

	// Сохраняем изменения в JSON-файл
	err := utils.SaveFile(config.TasksFilePath, newTasks)
	if err != nil {
		return err
	}

	fmt.Println("Task was deleted")
	return nil
}
