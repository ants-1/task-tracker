# Task Tracker

Task tracker is a simple command line interface (CLI) to track what you need to do, what you've done and what you're currently working on.


## Features

- List tasks
- Add new task
- Update a task description
- Delete task
- Mark task status
- Filter tasks by status


## Run Locally

Clone the project

```bash
  git clone https://github.com/ants-1/task-tracker.git
```

Go to the project directory

```bash
  cd task-tracker
```

Build the application

```bash
  go build -o ./task-tracker
```


## Usage/Examples

```go
- Adding a new task
./task-tracker add "Buy groceries"
- Output: Task added sucessfully!

- Updating and deleting tasks
./task-tracker update 1 "Buy groceries and cook dinner"
./task-tracker delete 1

- Marking a task as todo, in progress or done
./task-tracker mark-in-progress 1
./task-tracker mark-done 1
./task-tracker todo 1

- Listing all tasks
./task-tracker list

- Listing tasks by status
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```

