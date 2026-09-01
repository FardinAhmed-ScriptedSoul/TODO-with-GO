# TODO with Go

A lightweight command-line Todo manager built in Go. It lets you create, update, delete, and track tasks from the terminal and persists them in a local JSON file.

## Features

- Add new todo items
- Delete items by index
- Edit an existing todo title
- Toggle task completion status
- View all tasks in a formatted table
- Persist data to `todos.json`
- Load previous tasks automatically when the app starts
- Display completion timestamps for completed tasks

## Tech Stack

### Standard library
- `flag` — parses command-line flags
- `fmt` — prints messages and formats data
- `os` — file I/O and console output
- `encoding/json` — reads and writes JSON data
- `strconv` — converts indices to strings for table output
- `time` — stores creation and completion timestamps

### External library
- `github.com/aquasecurity/table` — renders the todo list in a clean terminal table

## Project Structure

- `main.go` — application entry point
- `todo.go` — todo model and table rendering logic
- `command.go` — CLI flag parsing and command execution
- `storage.go` — JSON save/load logic
- `todos.json` — local task storage file

## Getting Started

### Prerequisites

- Go 1.22 or newer
- Git (optional, for version control)

### Install dependencies

```bash
go mod tidy
```

### Run the app

```bash
go run .
```

## Command Documentation

The app uses command-line flags to perform actions.

### 1. List all todos

```bash
go run . -list
```

Example output:

```text
┌───┬─────────────────────┬───────────┬─────────────────────┬─────────────────────┐
│ # │ Title               │ Completed │     Created At      │    Completed At     │
├───┼─────────────────────┼───────────┼─────────────────────┼─────────────────────┤
│ 0 │ Buy groceries       │ ✅        │ 2026-09-02 10:30:15 │ 2026-09-02 11:00:10 │
│ 1 │ Walk the dog        │ ❌        │ 2026-09-02 12:00:00 │                     │
└───┴─────────────────────┴───────────┴─────────────────────┴─────────────────────┘
```

### 2. Add a new todo

```bash
go run . -add "Buy groceries"
```

This appends a new task to the list and saves it to `todos.json`.

### 3. Delete a todo by index

```bash
go run . -del 0
```

This removes the task at index `0`.

### 4. Toggle completion status

```bash
go run . -toggle 0
```

If the task is incomplete, it becomes completed and a completion timestamp is stored. If it is already complete, it is marked as incomplete and the completion timestamp is cleared.

### 5. Edit a todo title

```bash
go run . -edit "0:Groceries"
```

This updates the title of the task at index `0`.

> Note: the current implementation parses the value using a single word after the colon, so values with spaces should be passed in a format that matches the parser behavior exactly.

## Storage Behavior

The app automatically reads from and writes to `todos.json` in the project root.

Example `todos.json`:

```json
[
    [
        {
            "Title": "Buy groceries",
            "Completed": false,
            "CreatedAt": "2026-09-02T10:30:15Z",
            "CompletedAt": null
        }
    ]
]
```

## Example Workflow

```bash
go run . -add "Buy groceries"
go run . -add "Walk the dog"
go run . -list
go run . -toggle 0
go run . -edit "1:Walk the dog in the park"
go run . -del 0
go run . -list
```

## Notes

- The app stores all tasks in a JSON array structure.
- The list view is table-based and designed for terminal-friendly output.
- The app’s command behavior is driven by the `flag` package, which means all actions are invoked as CLI flags.
