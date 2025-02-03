package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var tasksdb []Task = make([]Task, 0)

const filePath string = "tasks.json"

func initDb() {
	if _, err := os.Create(filePath); err != nil {
		log.Fatalf("failed on create file tasks.json: %s", err)
	}

	if err := storeData(); err != nil {
		log.Fatal(err)
	}
}

func initProgram() {
	if _, err := os.Stat(filePath); err != nil {
		initDb()
		return
	}

	if err := loadData(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initProgram()

	flag.Parse()
	if flag.NArg() == 0 {
		// flag.Usage() -> TODO: implementar função para dar um help no comando
		os.Exit(1)
	}

	action := flag.Arg(0)
	param1 := flag.Arg(1)
	param2 := flag.Arg(2)

	switch action {
	case "add":
		addTask(param1)
	case "update":
		id, _ := strconv.Atoi(param1)
		updateTask(id, param2)
	case "delete":
		id, _ := strconv.Atoi(param1)
		deleteTask(id)
	case "mark-in-progress":
		id, _ := strconv.Atoi(param1)
		markTask(id, "in-progress")
	case "mark-done":
		id, _ := strconv.Atoi(param1)
		markTask(id, "done")
	case "list":
		listTask(param1)
	default:
		fmt.Printf("Command not found")
		os.Exit(1)
	}
}

func addTask(desc string) error {
	newTask := Task{
		ID:          len(tasksdb) + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasksdb = append(tasksdb, newTask)

	if err := storeData(); err != nil {
		return fmt.Errorf("failed on store data: %s", err)
	}

	fmt.Printf("Task added successfully (ID: %d)", newTask.ID)

	return nil
}

func updateTask(id int, desc string) error {
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

	task.Description = desc
	task.UpdatedAt = time.Now()

	tasksdb[idx] = task

	if err := storeData(); err != nil {
		return fmt.Errorf("failed on store data: %s", err)
	}

	fmt.Printf("Task updated successfully (ID: %d)", id)

	return nil
}

func deleteTask(id int) error {
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

	if err := storeData(); err != nil {
		return fmt.Errorf("failed on store date: %s", err)
	}

	return nil
}

func markTask(id int, status string) error {
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

	if err := storeData(); err != nil {
		return fmt.Errorf("failed on store date: %s", err)
	}

	return nil
}

func listTask(status string) error {
	mapStatus := map[string]string{
		"done":        "done",
		"todo":        "todo",
		"in-progress": "in-progress",
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

func storeData() error {
	data, err := json.Marshal(tasksdb)
	if err != nil {
		return fmt.Errorf("failed on marshal tasks: %s", err)
	}
	err = os.WriteFile("./tasks.json", data, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed on write file: %s", err)
	}

	return nil
}

func loadData() error {
	fmt.Println("loading data...")
	data, err := os.ReadFile("./tasks.json")
	if err != nil {
		return fmt.Errorf("failed on read file: %s", err)
	}

	err = json.Unmarshal(data, &tasksdb)
	if err != nil {
		return fmt.Errorf("failed on unmarshal: %s", err)
	}

	return nil
}
