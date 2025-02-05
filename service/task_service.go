package service

import (
	"fmt"
	"log"
	"time"

	"github.com/silas-ss/task-tracker-cli/model"
	"github.com/silas-ss/task-tracker-cli/repository"
)

type TaskService struct {
	TaskRepo repository.TaskRepository
}

const (
	StatusTodo = "todo"
	StatusInProgress = "in-progress"
	StatusDone = "done"
)

func (ts *TaskService) AddTask(desc string) error {
	newTask := model.Task{
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newTask, err := ts.TaskRepo.CreateTask(newTask)
	if err != nil {
		return fmt.Errorf("failed on add task. error: %s", err)
	}
	
	fmt.Printf("Task added successfully (ID: %d)", newTask.ID)

	return nil
}

func (ts *TaskService) UpdateTask(id int, desc string) error {
	task, err := ts.TaskRepo.FindTaskById(id)
	if err != nil {
		log.Fatalf("error on update task. err: %s", err)
	}

	task.Description = desc
	task.UpdatedAt = time.Now()

	if _, err := ts.TaskRepo.UpdateTask(task); err != nil {
		log.Fatalf("error on update task. error: %s", err)
	}

	fmt.Printf("Task updated successfully (ID: %d)", id)

	return nil
}

func (ts *TaskService) DeleteTask(id int) error {
	return ts.TaskRepo.DeleteById(id)
}

func (ts *TaskService) MarkTaskWithStatus(id int, status string) error {
	task, err := ts.TaskRepo.FindTaskById(id)
	if err != nil {
		log.Fatalf("error on update task. err: %s", err)
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	if _, err := ts.TaskRepo.UpdateTask(task); err != nil {
		return fmt.Errorf("failed on store date: %s", err)
	}

	return nil
}

func (ts *TaskService) ListTask(status string) error {
	tasks, err := ts.TaskRepo.FindByStatus(status)
	if err != nil {
		return fmt.Errorf("error on list tasks. error: %s", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	for i := 0; i < len(tasks); i++ {
		task := tasks[i]

		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Status: %s\n", task.Status)
		fmt.Printf("CreatedAt: %s\n", task.CreatedAt.Format("2006-02-01 15:04:05"))
		fmt.Printf("UpdatedAt: %s\n", task.UpdatedAt.Format("2006-02-01 15:04:05"))
		fmt.Println("--------------------------------------------------------")
	}

	return nil
}