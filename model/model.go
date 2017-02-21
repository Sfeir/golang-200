package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// TaskStatus is the current processing status of a task
type TaskStatus int

const (
	// StatusTodo is used for incomplete tasks
	StatusTodo TaskStatus = iota
	// StatusInProgress is used for tasks in progress
	StatusInProgress
	// StatusDone is used for completed tasks
	StatusDone
)

// TaskPriority is the priority of a task
type TaskPriority int

const (
	// PriorityMinor is used for task with a lower priority
	PriorityMinor TaskPriority = iota
	// PriorityMedium is used for task with medium priority
	PriorityMedium
	// PriorityHigh is used for task with high priority
	PriorityHigh
)

// Task is the structure to define a task to be done
type Task struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty" `
	Title        string        `json:"title" bson:"title"`
	Description  string        `json:"description" bson:"description"`
	Status       TaskStatus    `json:"status" bson:"status"`
	Priority     TaskPriority  `json:"priority" bson:"priority"`
	CreationDate time.Time     `json:"creationDate" bson:"creationDate"`
	DueDate      time.Time     `json:"dueDate" bson:"dueDate"`
}

// GetID returns the ID of a Task as a string
func (s *Task) GetID() string {
	return s.ID.Hex()
}
