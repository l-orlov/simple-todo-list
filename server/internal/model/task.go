package model

import (
	"time"

	"github.com/google/uuid"
)

// TaskToCreate - данные для создания таски
type TaskToCreate struct {
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}

// TaskToUpdate - данные для обновления таски
type TaskToUpdate struct {
	ID     uuid.UUID  `json:"id"`
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}

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
