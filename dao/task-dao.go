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

	// GetByTitle returns all tasks by title
	GetByTitle(title string) ([]model.Task, error)

	// GetByStatus returns all tasks by status
	GetByStatus(status model.TaskStatus) ([]model.Task, error)

	// GetByStatusAndPriority returns all tasks by status and priority
	GetByStatusAndPriority(status model.TaskStatus, priority model.TaskPriority) ([]model.Task, error)

	// Save saves the task
	Save(task *model.Task) error

	// Upsert updates or creates a task
	Upsert(task *model.Task) (bool, error)

	// Delete deletes a tasks by its ID
	Delete(ID string) error
}
