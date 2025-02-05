package repository

import (
	"encoding/json"
	"fmt"

	"github.com/silas-ss/task-tracker-cli/db"
	"github.com/silas-ss/task-tracker-cli/model"
)

type TaskRepository struct {
	Fs db.FileSystem
}

func NewTaskRepository() TaskRepository {
	fs := db.FileSystem{}
	fs.InitProgram()
	return TaskRepository{Fs: fs}
}

func (tr *TaskRepository) FindAll() ([]model.Task, error) {
	tasks := []model.Task{}
	data, err := tr.Fs.LoadData()
	if err != nil {
		return tasks, err
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return tasks, fmt.Errorf("error on unmarshal tasks. error: %s", err)
	}

	return tasks, nil
}

func (tr *TaskRepository) FindByStatus(status string) ([]model.Task, error) {
	tasks, err := tr.FindAll()
	if err != nil {
		return tasks, err
	}

	// when no filter
	if len(status) == 0 {
		return tasks, nil
	}

	filtered := []model.Task{}
	for i := 0; i < len(tasks); i++ {
		t := tasks[i]
		if t.Status == status {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}

func (tr *TaskRepository) FindTaskById(id int) (model.Task, error) {
	tasks, err := tr.FindAll()
	if err != nil {
		return model.Task{}, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return model.Task{}, fmt.Errorf("task with id: %d not found", id)
}

func (tr *TaskRepository) CreateTask(task model.Task) (model.Task, error) {
	tasks, err := tr.FindAll()
	if err != nil {
		return model.Task{}, err
	}

	// set new id
	task.ID = len(tasks) + 1

	tasks = append(tasks, task)
	if err := tr.Fs.StoreData(tasks); err != nil {
		return model.Task{}, fmt.Errorf("error on store data. error: %s", err)
	}

	return task, nil
}

func (tr *TaskRepository) UpdateTask(task model.Task) (model.Task, error) {
	tasks, err := tr.FindAll()
	if err != nil {
		return model.Task{}, err
	}

	idx := -1
	for i := 0; i < len(tasks); i++ {
		t := tasks[i]
		if t.ID == task.ID {
			idx = i
			break
		}
	}
	if idx < 0 {
		return model.Task{}, fmt.Errorf("task with id: %d not found", task.ID)
	}

	tasks[idx] = task

	if err := tr.Fs.StoreData(tasks); err != nil {
		return model.Task{}, fmt.Errorf("failed on store data: %s", err)
	}

	return task, nil
}

func (tr *TaskRepository) DeleteById(id int) error {
	tasks, err := tr.FindAll()
	if err != nil {
		return err
	}

	idx := -1
	for i := 0; i < len(tasks); i++ {
		t := tasks[i]
		if t.ID == id {
			idx = i
			break
		}
	}

	if idx < 0 {
		return fmt.Errorf("task not found with id: %d", id)
	}

	tasks = append(tasks[:idx], tasks[idx+1:]...)

	if err := tr.Fs.StoreData(tasks); err != nil {
		return fmt.Errorf("error on store data. error: %s", err)
	}

	return nil
}