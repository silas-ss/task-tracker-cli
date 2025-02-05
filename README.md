# Task Tracker CLI

This project is a solution for the [Task Tracker](https://roadmap.sh/projects/task-tracker) challenge from roadmap.sh.

## How to run

Clone this repository and run the following command inside directory:

```bash
go build -o task-tracker
```

Running this application

```bash
# Add task
./task-tracker add "Buy groceries"

# List tasks
./task-tracker list

# List tasks done
./task-tracker list done

# List tasks todo
./task-tracker list todo

# List tasks in-progress
./task-tracker list in-progress

# Update task description
./task-tracker update 1 "Buy groceries and cheese"

# Mark task as in-progress
./task-trakcer mark-in-progress 1

# Mark task as done
./task-tracker mark-done 1

# Delete task
./task-tracker delete 1
```
