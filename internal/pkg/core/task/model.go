package task

import (
	"time"

	"github.com/google/uuid"
)

// Task describes base task entity
type Task struct {
	ID          string
	Description string
	Summary     string
	Created     time.Time
	Updated     time.Time
	Closed      time.Time
}

// NewTask creates new instance of task
func NewTask(summary, description string) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Summary:     summary,
		Description: description,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
}
