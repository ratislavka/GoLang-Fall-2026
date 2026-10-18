# SIS Project - Checkpoint 1

## Overview
Simple in-memory Go application that creates, stores, and processes text analysis tasks. Each task takes a string input and calculates:
- Word count
- Character count (without spaces)
- Vowel count

## Structure
- `main.go` - runs a quick example of creating and processing a task
- `task/task.go` - `Task` and `Result` structs, input validation, and status transitions
- `task/processor.go` - `Processor` interface and `TextProcessor` implementation
- `task/manager.go` - `Manager` struct that stores tasks in a map in memory
- `task/task_test.go` - unit tests

## Task Lifecycle
Tasks go through 3 states:
`Pending` -> `Processing` -> `Completed`

- `Pending`: task is created and validated, waiting to be processed.
- `Processing`: task is currently running.
- `Completed`: processing is finished and the result is saved.

Invalid transitions (like trying to process a task that is already `Completed` or jumping straight from `Pending` to `Completed`) return an error.

## How to Run
Run the main file:
```bash
go run .
```

Run checks and tests:
```bash
gofmt -l .
go vet ./...
go test -v ./...
```