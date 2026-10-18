package task

import (
	"errors"
	"strings"
)

const (
	StatusPending    = "Pending"
	StatusProcessing = "Processing"
	StatusCompleted  = "Completed"
)

type Result struct {
	WordCount  int
	CharCount  int
	VowelCount int
}

type Task struct {
	ID     int
	Input  string
	Status string
	Result Result
}

func (t Task) Validate() error {
	if strings.TrimSpace(t.Input) == "" {
		return errors.New("input text cannot be empty")
	}
	return nil
}

// pointer receiver here cause we need to actually change the status field on the struct
func (t *Task) UpdateStatus(newStatus string) error {
	if t.Status == StatusPending && newStatus == StatusProcessing {
		t.Status = newStatus
		return nil
	}
	if t.Status == StatusProcessing && newStatus == StatusCompleted {
		t.Status = newStatus
		return nil
	}
	return errors.New("invalid status transition from " + t.Status + " to " + newStatus)
}
