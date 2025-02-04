package main

import (
	"fmt"
	"log"
	"time"
)

type TaskService struct {
	TaskRepo TaskRepository
}

func (ts *TaskService) addTask(desc string) error {
	newTask := Task{
		ID:          len(tasksdb) + 1,
		Description: desc,
		Status:      statusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasksdb = append(tasksdb, newTask)

	if err := ts.TaskRepo.Fs.storeData(); err != nil {
		return fmt.Errorf("failed on store data: %s", err)
	}

	fmt.Printf("Task added successfully (ID: %d)", newTask.ID)

	return nil
}



func (ts *TaskService) updateTask(id int, desc string) error {
	task, err := ts.TaskRepo.findTaskById(id)
	if err != nil {
		log.Fatalf("error on update task. err: %s", err)
	}

	task.Description = desc
	task.UpdatedAt = time.Now()

	if _, err := ts.TaskRepo.saveTask(task); err != nil {
		log.Fatalf("error on update task. error: %s", err)
	}

	fmt.Printf("Task updated successfully (ID: %d)", id)

	return nil
}

func (ts *TaskService) deleteTask(id int) error {
	idx := -1
	for i := 0; i < len(tasksdb); i++ {
		t := tasksdb[i]
		if t.ID == id {
			idx = i
			break
		}
	}

	if idx < 0 {
		fmt.Printf("task not found with id: %d\n", id)
		return fmt.Errorf("task not found with id: %d", id)
	}

	tasksdb = append(tasksdb[:idx], tasksdb[idx+1:]...)

	if err := ts.TaskRepo.Fs.storeData(); err != nil {
		return fmt.Errorf("failed on store date: %s", err)
	}

	return nil
}

func (ts *TaskService) markTask(id int, status string) error {
	var task Task
	idx := -1
	for i := 0; i < len(tasksdb); i++ {
		t := tasksdb[i]
		if t.ID == id {
			task = t
			idx = i
			break
		}
	}

	if idx < 0 {
		fmt.Printf("task not found with id: %d\n", id)
		return fmt.Errorf("task not found with id: %d", id)
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	tasksdb[idx] = task

	if err := ts.TaskRepo.Fs.storeData(); err != nil {
		return fmt.Errorf("failed on store date: %s", err)
	}

	return nil
}

func (ts *TaskService) listTask(status string) error {
	mapStatus := map[string]string{
		statusDone:        statusDone,
		statusTodo:        statusTodo,
		statusInProgress:  statusInProgress,
		"":            "all",
	}

	criteria := mapStatus[status]

	if len(tasksdb) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	for i := 0; i < len(tasksdb); i++ {
		task := tasksdb[i]

		if criteria != "all" && criteria != task.Status {
			continue
		}

		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Status: %s\n", task.Status)
		fmt.Printf("CreatedAt: %s\n", task.CreatedAt.Format("2006-02-01 15:04:05"))
		fmt.Printf("UpdatedAt: %s\n", task.UpdatedAt.Format("2006-02-01 15:04:05"))
		fmt.Println("--------------------------------------------------------")
	}

	return nil
}