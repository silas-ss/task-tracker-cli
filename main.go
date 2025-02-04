package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	filePath string = "./tasks.json"

	cmdAdd string = "add"
	cmdList string = "list"
	cmdUpdate string = "update"
	cmdDelete string = "delete"
	
	cmdMarkDone string = "mark-done"
	cmdMarkInProgress string = "mark-in-progress"

	statusTodo = "todo"
	statusInProgress = "in-progress"
	statusDone = "done"
)

func main() {
	fs := FileSystem{}
	ts := TaskService{
		TaskRepo: TaskRepository{
			Fs: fs,
		},
	}

	fs.initProgram()

	flag.Parse()
	if flag.NArg() == 0 {
		// flag.Usage() -> TODO: implementar função para dar um help no comando
		os.Exit(1)
	}

	action := flag.Arg(0)
	param1 := flag.Arg(1)
	param2 := flag.Arg(2)

	handleCommand(ts, action, param1, param2)
}

func handleCommand(ts TaskService, action string, params ...string) {
	param1 := params[0]
	param2 := params[1]

	switch action {
	case cmdAdd:
		ts.addTask(param1)
	case cmdUpdate:
		id, _ := strconv.Atoi(param1)
		ts.updateTask(id, param2)
	case cmdDelete:
		id, _ := strconv.Atoi(param1)
		ts.deleteTask(id)
	case cmdList:
		ts.listTask(param1)
	case cmdMarkInProgress:
		id, _ := strconv.Atoi(param1)
		ts.markTask(id, statusInProgress)
	case cmdMarkDone:
		id, _ := strconv.Atoi(param1)
		ts.markTask(id, statusDone)
	default:
		fmt.Printf("Command not found")
		os.Exit(1)
	}
}


