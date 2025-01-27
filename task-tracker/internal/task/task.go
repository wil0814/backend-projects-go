package task

import (
	"backend-products-go/task-tracker/internal/file"
	"backend-products-go/task-tracker/internal/model"
	"fmt"
	"time"
)

func AddTask(description string) (int64, error) {
	tasks, err := file.ReadTasksFromFile()
	if err != nil {
		return 0, err
	}

	var newID int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newID = lastTask.ID + 1
	} else {
		newID = 1
	}

	tasks = append(tasks, *model.NewTask(newID, description))

	err = file.WriteTasksToFile(tasks)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func UpdateTask(id int64, description string) error {
	tasks, err := file.ReadTasksFromFile()
	if err != nil {
		return err
	}

	tasksExist := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			tasksExist = true
			break
		}
	}

	if !tasksExist {
		return fmt.Errorf("task with id %d not found", id)
	}

	return file.WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := file.ReadTasksFromFile()
	if err != nil {
		return err
	}

	tasksExist := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			tasksExist = true
			break
		}
	}

	if !tasksExist {
		return fmt.Errorf("task with id %d not found", id)
	}

	return file.WriteTasksToFile(tasks)
}

func MarkTaskAsInProgress(id int64) error {
	tasks, err := file.ReadTasksFromFile()
	if err != nil {
		return err
	}

	tasksExist := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = model.TaskStatusInProgress
			tasks[i].UpdatedAt = time.Now()
			tasksExist = true
			break
		}
	}

	if !tasksExist {
		return fmt.Errorf("task with id %d not found", id)
	}

	return file.WriteTasksToFile(tasks)
}

func MarkTaskAsDone(id int64) error {
	tasks, err := file.ReadTasksFromFile()
	if err != nil {
		return err
	}

	tasksExist := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = model.TaskStatusDone
			tasks[i].UpdatedAt = time.Now()
			tasksExist = true
			break
		}
	}

	if !tasksExist {
		return fmt.Errorf("task with id %d not found", id)
	}

	return file.WriteTasksToFile(tasks)
}

func ListTasks(status *model.TaskStatus) ([]model.Task, error) {
	tasks, err := file.ReadTasksFromFile()
	if err != nil {
		return nil, err
	}

	if status == nil {
		return tasks, nil
	}

	var filteredTasks []model.Task

	for _, task := range tasks {
		if task.Status == *status {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}
