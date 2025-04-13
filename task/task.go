package task

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var file = "tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &tasks)
	}
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}
