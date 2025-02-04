package main

import "fmt"

type TaskRepository struct {
	Fs FileSystem
}

var tasksdb []Task = make([]Task, 0)

func (tr *TaskRepository) findTaskById(id int) (Task, error) {
	for i := 0; i < len(tasksdb); i++ {
		t := tasksdb[i]
		if t.ID == id {
			return t, nil
		}
	}

	return Task{}, fmt.Errorf("task with id: %d not found", id)
}

func (tr *TaskRepository) saveTask(task Task) (Task, error) {
	idx := -1
	for i := 0; i < len(tasksdb); i++ {
		t := tasksdb[i]
		if t.ID == task.ID {
			idx = i
			break
		}
	}
	if idx < 0 {
		tasksdb = append(tasksdb, task)
		idx = len(tasksdb) - 1
	}

	tasksdb[idx] = task

	if err := tr.Fs.storeData(); err != nil {
		return Task{}, fmt.Errorf("failed on store data: %s", err)
	}

	return task, nil
}