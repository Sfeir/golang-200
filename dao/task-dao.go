package dao

import (
	"github.com/Sfeir/golang-200/model"
)

const (
	// NoPaging used with skip, limit parameters
	NoPaging = -1
)

// TaskDAO is the DAO interface to work with tasks
type TaskDAO interface {

	// GetByID returns a task by its ID
	GetByID(ID string) (*model.Task, error)

	// GetAll returns all tasks with paging capability
	GetAll(start, end int) ([]model.Task, error)

	// GetByName returns all tasks by name
	GetByName(name string) ([]model.Task, error)

	// GetByType returns all tasks by type
	GetByType(taskType string) ([]model.Task, error)

	// GetByTypeAndScore returns all tasks by type and score greater than parameter
	GetByTypeAndScore(taskType string, score uint8) ([]model.Task, error)

	// Save saves the task
	Save(task *model.Task) error

	// Upsert updates or creates a task
	Upsert(ID string, task *model.Task) (bool, error)

	// Delete deletes a tasks by its ID
	Delete(ID string) error
}
