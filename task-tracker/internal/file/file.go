package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wil0814/backend-projects-go/task-tracker/internal/model"
)

const fileName = "tasks.json"

func getFilePath() string {
	if filePath := os.Getenv("TASKS_FILE_PATH"); filePath != "" {
		return filepath.Join(filePath, fileName)
	}

	return fileName
}

func ensureFileInitialized(filePath string) error {
	// Check if the file exists and create it if it doesn't
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File not found. Creating a new one...")
		return os.WriteFile(filePath, []byte("[]"), 0644)
	}
	return nil
}

func ReadTasksFromFile() ([]model.Task, error) {
	filePath := getFilePath()
	var tasks []model.Task

	if err := ensureFileInitialized(filePath); err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []model.Task) error {
	filePath := getFilePath()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}
