package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "inprogress"
	StatusDone       = "done"
	PriorityMinor    = iota
	PriorityMedium
	PriorityHigh
	PriorityCritical
)

// Task is the structure to define a task to be done
type Task struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty" `
	Title        string        `json:"title" bson:"title"`
	Description  string        `json:"description" bson:"description"`
	Status       string        `json:"status" bson:"status"`
	Priority     int           `json:"priority" bson:"priority"`
	CreationDate time.Time     `json:"creationDate" bson:"creationDate"`
}

// GetID returns the ID of a Task as a string
func (s *Task) GetID() string {
	return s.ID.Hex()
}
