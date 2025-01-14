package commands

import (
	"fmt"
	"task-manager-cli/models"
	"task-manager-cli/utils"
	"time"
)

func UpdateTask(tasks []models.Task, id int, newDescription string) error {
	var isFounded bool
	for i, task := range tasks {
		if task.ID == id {
			isFounded = true
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			break
		}
	}
	if !isFounded {
		fmt.Println("Task not found")
	}

	// Сохраняем изменения в JSON-файл
	err := utils.SaveFile("tasks.json", tasks)

	if err != nil {
		return err
	}

	fmt.Println("Task was updated")
	return nil
}
