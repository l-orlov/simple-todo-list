package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Status    TaskStatus `json:"status" db:"status"`
	Title     string     `json:"title" db:"title"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
}

// DbTable returns DB table name
func (Task) DbTable() string {
	return "tasks"
}

type TaskStatus int32

const (
	TaskStatusToDo       = TaskStatus(1)
	TaskStatusInProgress = TaskStatus(2)
	TaskStatusDone       = TaskStatus(3)
	TaskStatusDeleted    = TaskStatus(4)
)
