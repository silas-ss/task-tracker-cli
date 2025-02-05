package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/silas-ss/task-tracker-cli/repository"
	"github.com/silas-ss/task-tracker-cli/service"
)

const (
	cmdAdd string = "add"
	cmdList string = "list"
	cmdUpdate string = "update"
	cmdDelete string = "delete"
	
	cmdMarkDone string = "mark-done"
	cmdMarkInProgress string = "mark-in-progress"
)

func main() {
	ts := service.TaskService{
		TaskRepo: repository.NewTaskRepository(),
	}

	flag.Parse()
	if flag.NArg() == 0 {
		os.Exit(1)
	}

	action := flag.Arg(0)
	param1 := flag.Arg(1)
	param2 := flag.Arg(2)

	handleCommand(ts, action, param1, param2)
}

func handleCommand(ts service.TaskService, action string, params ...string) {
	param1 := params[0]
	param2 := params[1]

	switch action {
	case cmdAdd:
		ts.AddTask(param1)
	case cmdUpdate:
		id, _ := strconv.Atoi(param1)
		ts.UpdateTask(id, param2)
	case cmdDelete:
		id, _ := strconv.Atoi(param1)
		ts.DeleteTask(id)
	case cmdList:
		ts.ListTask(param1)
	case cmdMarkInProgress:
		id, _ := strconv.Atoi(param1)
		ts.MarkTaskWithStatus(id, service.StatusInProgress)
	case cmdMarkDone:
		id, _ := strconv.Atoi(param1)
		ts.MarkTaskWithStatus(id, service.StatusDone)
	default:
		fmt.Printf("Command not found")
		os.Exit(1)
	}
}


