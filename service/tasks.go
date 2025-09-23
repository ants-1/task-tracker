package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"` // todo, in-progress, done
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const fileName = "tasks.json"

func CreateTaskFile() {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("task.json does not exist")

		err = os.WriteFile("tasks.json", []byte("[]"), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		fmt.Println("tasks.json created successfully")
	}
	defer file.Close()
}

func AddTask(task Task) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if len(tasks) == 0 {
		task.ID = 1
	} else {
		task.ID = tasks[len(tasks)-1].ID + 1
	}
	task.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	tasks = append(tasks, task)

	newData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	os.WriteFile(fileName, newData, 0644)
	fmt.Println("New task added:", task)
	return os.WriteFile(fileName, newData, 0644)
}

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error loading tasks.json file:", err)
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Error parsing tasks.json:", err)
		return nil, err
	}

	return tasks, nil
}

func FilteredTasks(status string) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	var filteredList []Task

	for _, t := range tasks {
		if t.Status == status {
			filteredList = append(filteredList, t)
		}
	}
	return filteredList, nil
}

func EditTask(id int, newDescription string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			if newDescription != "" {
				tasks[i].Description = newDescription
			}
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with ID %d not found", id)
	}

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func EditTaskStatus(id int, newStatus string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			if newStatus != "" {
				tasks[i].Status = newStatus
			}
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Task with ID %d not found", id)
	}

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return fmt.Errorf("Task with ID %d not found", id)
	}

	data, err := json.MarshalIndent(newTasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
