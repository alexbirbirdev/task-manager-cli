package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"task-manager-cli/models"
)

func LoadTasks(filePath string, tasks interface{}) error {
	_, err := os.Stat(filePath)

	// Проверяем, существует ли файл. Если нет, создаем
	if os.IsNotExist(err) {
		_, err := os.Create("tasks.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		emptyTasks := []models.Task{}
		SaveFile("tasks.json", emptyTasks)

		// Если файл существует, работаем с ним
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		return decoder.Decode(tasks)
	}
	return nil
}

// Сохранение файла
func SaveFile(filePath string, tasks interface{}) error {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(tasks)
}
