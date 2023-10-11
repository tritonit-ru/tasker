package task

import (
	"context"
	"time"
)

// Provider describes interface to getting task information
type Provider interface {
	ID() string
	Description() string
	Summary() string
	Created() time.Time
	Updated() time.Time
	Closed() time.Time
}

// Repository describes interface for task operations
type Repository interface {
	Create(ctx context.Context, task Provider) (taskID int64, err error)
	Update(ctx context.Context, task Provider) (err error)
	GetByID(ctx context.Context, id string) (task Provider, err error)
	FetchAll() ([]Provider, error)
}

// Service describes interface of task manager
type Service interface {
	CreateTask(ctx context.Context, summary, description string) (task *Task, err error)
	GetTask(ctx context.Context, taskID string) (task *Task, err error)
	FetchTasks(ctx context.Context) (tasks []*Task, err error)
	CloseTask(ctx context.Context, taskID string) (err error)
}
