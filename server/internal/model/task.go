package model

import (
	"time"

	"github.com/google/uuid"
)

// Task - таска
type Task struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	UserID    uuid.UUID  `json:"-" db:"user_id"`
	Title     string     `json:"title" db:"title"`
	Status    TaskStatus `json:"status" db:"status"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
}

// DbTable returns DB table name
func (Task) DbTable() string {
	return "tasks"
}

// TaskStatus - статус таски
type TaskStatus int32

const (
	TaskStatusToDo       = TaskStatus(1)
	TaskStatusInProgress = TaskStatus(2)
	TaskStatusDone       = TaskStatus(3)
	TaskStatusDeleted    = TaskStatus(4)
)
