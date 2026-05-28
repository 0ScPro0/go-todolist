package domain

import "time"

// TaskStatus type
type TaskStatus int

const (
	NEW TaskStatus = iota
	IN_PROGRESS
	DONE
)

func (s TaskStatus) String() string {
	return [...]string{"new", "in_progress", "done"}[s]
}

// Domain entity Task
// Task has required fields Name and Status
// and optional field Description
type Task struct {
	ID      int `json:"id"`
	Version int `json:"version"`

	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt 	time.Time  `json:"created_at"`
}