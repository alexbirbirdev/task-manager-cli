package commands

import (
	"fmt"
	"task-manager-cli/models"
	"task-manager-cli/utils"
	"time"
)

func AddTask(description string, tasks []models.Task) error {
	// Инициализируем новую задачу
	newTask := models.Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	// Добавляем в слайс новую задачу
	tasks = append(tasks, newTask)

	// Сохраняем изменения в JSON-файл
	err := utils.SaveFile("tasks.json", tasks)
	if err != nil {
		return err
	}

	fmt.Println("Task was added")
	return nil
}
