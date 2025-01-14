package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:updatedAt"`
}

func main() {
	// Пробуем открыть файл
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [string]")
		os.Exit(1)
	}

	command := os.Args[1]
	// command := ""

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run add <title>")
			os.Exit(1)
		}
		description := os.Args[2]
		addTask(description, tasks)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run update <id> <title>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		description := os.Args[3]
		// fmt.Printf("Updating task with ID %v to title: %s\n", id, description)
		//Исполнение функционала изменения заголовка
		updateTask(tasks, id, description)

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
		var status string
		if len(os.Args) == 3 {
			status = os.Args[2]
		}
		listTasks(tasks, status)
	default:
		fmt.Printf("Invalid command: %s\n", command)
		os.Exit(1)
	}
}

func addTask(description string, tasks []Task) ([]Task, error) {
	// Инициализируем новую задачу
	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	// Добавляем в слайс новую задачу
	tasks = append(tasks, newTask)

	// Сохраняем изменения в JSON-файл
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return tasks, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func listTasks(tasks []Task, status string) {
	c := 0
	for _, task := range tasks {
		if task.Status == status || status == "" {
			fmt.Printf("ID: %v\n Description: %s\n Status: %s\n CreatedAt: %s\n UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
			c++
		}
	}
	if c == 0 {
		fmt.Println("No tasks with status", status)
	}
}

func updateTask(tasks []Task, id int, newDescription string) ([]Task, error) {
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
		return tasks, fmt.Errorf("Task with ID %v not found", id)
	}

	// Сохраняем изменения в JSON-файл
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return tasks, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(tasks)

	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
